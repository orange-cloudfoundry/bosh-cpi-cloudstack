package action

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
	"github.com/xanzy/go-cloudstack/cloudstack"
	"strings"
	"time"
)

const (
	pvDriverErr = "VM which requires PV drivers to be installed"
)

func (a CPI) CreateVM(
	agentID apiv1.AgentID, stemcellCID apiv1.StemcellCID,
	cloudProps apiv1.VMCloudProps, networks apiv1.Networks,
	associatedDiskCIDs []apiv1.DiskCID, env apiv1.VMEnv) (apiv1.VMCID, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.CreateVm)

	var resProps ResourceCloudProperties
	err := cloudProps.As(&resProps)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Cannot create vm")
	}
	vmName := fmt.Sprintf("%s%s", config.VMPrefix, uuid.NewV4().String())

	userData := NewUserDataContents(vmName, a.config.Actions.Registry, networks)
	userDataRaw, _ := json.Marshal(userData)
	userDataStr := base64.StdEncoding.EncodeToString(userDataRaw)

	err = a.checkNetworkConfig(networks)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Error when creating vm")
	}

	defaultNetwork := a.findDefaultNetwork(networks)
	if defaultNetwork == nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Cannot found default network when creating vm")
	}

	if defaultNetwork.Type() == string(config.ManualNetwork) && defaultNetwork.IP() == "" {
		return apiv1.VMCID{}, bosherr.Errorf("Ip must be defined on a manual network")
	}

	var networkProps NetworkCloudProperties
	err = defaultNetwork.CloudProps().As(&networkProps)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Error when creating vm")
	}

	zoneId, err := a.findZoneId()
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Could not found zone when creating vm")
	}

	network, err := a.findNetworkByName(networkProps.Name, zoneId)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapErrorf(err, "Could not found network %s when creating vm", networkProps.Name)
	}

	serviceOffering, err := a.findServiceOfferingByName(resProps.ComputeOffering)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapErrorf(err, "Could not found compute offering %s when creating vm", resProps.ComputeOffering)
	}

	template, err := a.findTemplateByName(stemcellCID.AsString())
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapErrorf(err, "Could not found compute offering %s when creating vm", resProps.ComputeOffering)
	}

	a.logger.Info("create_vm", "Creating vm %s ...", vmName)
	deplParams := a.client.VirtualMachine.NewDeployVirtualMachineParams(serviceOffering.Id, template.Id, zoneId)
	deplParams.SetUserdata(userDataStr)
	deplParams.SetName(vmName)
	deplParams.SetKeypair(a.config.CloudStack.DefaultKeyName)
	deplParams.AddIptonetworklist(a.networkToMap(network, defaultNetwork.Type(), defaultNetwork.IP()))

	netParams, err := a.generateNetworksMap(networks, networkProps.Name, zoneId)
	if err != nil {
		return apiv1.VMCID{}, err
	}
	for _, netParam := range netParams {
		deplParams.AddIptonetworklist(netParam)
	}

	affinId, err := a.generateAffinityGroup(resProps, env)
	if err != nil {
		return apiv1.VMCID{}, err
	}
	if affinId != "" {
		deplParams.SetAffinitygroupids([]string{affinId})
	}

	if resProps.RootDiskSize > 0 {
		deplParams.SetRootdisksize(int64(resProps.RootDiskSize / 1024))
	}

	resp, err := a.client.VirtualMachine.DeployVirtualMachine(deplParams)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Error when creating vm")
	}

	if config.ToVmState(resp.State) != config.VmRunning {
		return apiv1.VMCID{}, a.destroyVmErrFallback(bosherr.Errorf("Vm is not running, actual state is %s", resp.State), resp.Id)
	}
	a.logger.Info("create_vm", "Finished creating vm %s .", vmName)

	a.logger.Info("create_vm", "Registering vm %s in registry...", vmName)
	vmCID := apiv1.NewVMCID(vmName)

	envSvc := a.regFactory.Create(vmCID)
	agentEnv := apiv1.NewAgentEnvFactory().ForVM(agentID, vmCID, networks, env, a.config.Actions.Agent)
	agentEnv.AttachSystemDisk("/dev/xvda")
	agentEnv.AttachEphemeralDisk("/dev/xvdb")

	err = envSvc.Update(agentEnv)
	if err != nil {
		return apiv1.VMCID{}, a.destroyVmErrFallback(bosherr.WrapError(err, "Error when creating vm"), resp.Id)
	}
	a.logger.Info("create_vm", "Finished registering vm %s in registry.", vmName)

	a.logger.Info("create_vm", "Creating vip(s) for vm %s ...", vmName)
	err = a.createVips(networks, resp.Id, zoneId, network)
	if err != nil {
		return apiv1.VMCID{}, a.destroyVmErrFallback(bosherr.WrapErrorf(err, "Could not create vips"), resp.Id)
	}
	a.logger.Info("create_vm", "Finished creating vip(s) for vm %s .", vmName)

	a.logger.Info("create_vm", "Creating ephemeral disk for vm %s ...", vmName)
	diskCid, err := a.createEphemeralDisk(resProps.EphemeralDiskSize, resProps.DiskCloudProperties, "")
	if err != nil {
		return apiv1.VMCID{}, a.destroyVmErrFallback(bosherr.WrapError(err, "Cannot create ephemeral disk when creating vm"), resp.Id)
	}
	a.logger.Info("create_vm", "Finished creating ephemeral disk for vm %s .", vmName)

	if len(resProps.Lbs) > 0 {
		a.logger.Info("create_vm", "Assigning vm %s to loadbalancers ...", vmName)
		err = a.setLoadBalancers(resProps.LBCloudProperties, resp.Id, zoneId, network.Id)
		if err != nil {
			return apiv1.VMCID{}, a.destroyVmErrFallback(bosherr.WrapError(err, "Cannot assign loadbalancers to vm"), resp.Id)
		}
		a.logger.Info("create_vm", "Finished assigning vm %s to loadbalancers.", vmName)
	}

	a.logger.Info("create_vm", "Attaching ephemeral disk for vm %s ...", vmName)
	err = a.attachEphemeralDisk(vmCID, diskCid)
	if err != nil {
		return apiv1.VMCID{}, a.destroyVmErrFallback(
			bosherr.WrapError(
				err,
				"Cannot attach ephemeral disk when creating vm"),
			resp.Id,
			func() {
				a.DeleteDisk(diskCid)
			},
		)
	}
	a.logger.Info("create_vm", "Finished attaching ephemeral disk for vm %s .", vmName)

	return vmCID, nil
}

func (a CPI) attachEphemeralDisk(vmCID apiv1.VMCID, diskCid apiv1.DiskCID) error {
	return retryable(5*time.Minute,
		func() error {
			return a.AttachDisk(vmCID, diskCid)
		},
		func(err error) bool {
			return strings.Contains(err.Error(), pvDriverErr)
		})
}

func (a CPI) destroyVmErrFallback(err error, vmId string, fs ...func()) error {
	a.deleteVMById(vmId)
	for _, fs := range fs {
		fs()
	}
	return err
}

func (a CPI) createVips(networks apiv1.Networks, vmId, zoneId string, defNetwork *cloudstack.Network) error {
	for _, network := range networks {
		if network.Type() != string(config.VipNetwork) {
			continue
		}
		err := a.createVip(network, vmId, zoneId, defNetwork)
		if err != nil {
			return bosherr.WrapErrorf(err, "Error when creating vip %s", network.IP())
		}
	}
	return nil
}

func (a CPI) createVip(network apiv1.Network, vmId, zoneId string, defNetwork *cloudstack.Network) error {
	if network.IP() == "" {
		return bosherr.Errorf("Vip must have ip defined")
	}

	publicIp, err := a.findPublicIpByIp(network.IP())
	if err != nil {
		return err
	}

	var networkProps NetworkCloudProperties
	err = network.CloudProps().As(&networkProps)
	if err != nil {
		return bosherr.WrapError(err, "Cannot get network properties")
	}

	networkCs := defNetwork
	if networkProps.Name != "" {
		networkCs, err = a.findNetworkByName(networkProps.Name, zoneId)
		if err != nil {
			return bosherr.WrapErrorf(err, "Could not found network %s", networkProps.Name)
		}
	}

	p := a.client.NAT.NewEnableStaticNatParams(publicIp.Id, vmId)
	p.SetNetworkid(networkCs.Id)

	return retryable(16*time.Minute,
		func() error {
			_, err := a.client.NAT.EnableStaticNat(p)
			return err
		},
		func(err error) bool {
			_, ok := err.(*json.SyntaxError)
			return ok
		})
}

func (a CPI) findDefaultNetwork(networks apiv1.Networks) apiv1.Network {
	if networks.Default().IP() != "" {
		return networks.Default()
	}
	for _, network := range networks {
		if network.Type() == string(config.ManualNetwork) {
			return network
		}
	}
	for _, network := range networks {
		if network.IsDynamic() {
			return network
		}
	}

	return nil
}

func (a CPI) checkNetworkConfig(networks apiv1.Networks) error {
	var nbManual int
	var nbDynamic int
	var nbVip int

	for _, network := range networks {
		if network.Type() == string(config.ManualNetwork) {
			nbManual++
		}
		if network.Type() == string(config.DynamicNetwork) {
			nbDynamic++
		}
		if network.Type() == string(config.VipNetwork) {
			nbVip++
		}
	}
	if nbVip > 1 {
		return bosherr.Errorf("Only 1 vip is supported")
	}
	if (nbDynamic + nbManual) == 0 {
		return bosherr.Errorf("It must have, at least, one dynamic or one manual network defined")
	}
	return nil
}

func (a CPI) generateAffinityGroup(resProps ResourceCloudProperties, env apiv1.VMEnv) (string, error) {
	if a.config.CloudStack.EnableAutoAntiAffinity {
		return a.generateAutoAffinityGroup(env)
	}
	if resProps.AffinityGroup == "" {
		return "", nil
	}
	affiType := resProps.AffinityGroupType
	if affiType == "" {
		affiType = "host anti-affinity"
	}
	affiId, err := a.findOrCreateAffinityGroup(resProps.AffinityGroup, affiType)
	if err != nil {
		return "", bosherr.WrapErrorf(
			err,
			"Could not find or create affinity group '%s' when creating vm",
			resProps.AffinityGroup)
	}
	return affiId, nil
}

func (a CPI) generateAutoAffinityGroup(env apiv1.VMEnv) (string, error) {
	vmEnv := NewVMEnv(env)
	name := fmt.Sprintf("%s-%s", a.ctx.DirectorUUID, vmEnv.Bosh.Group)
	affiId, err := a.findOrCreateAffinityGroup(name, "host anti-affinity")
	if err != nil {
		return "", bosherr.WrapErrorf(
			err,
			"Could not find or create affinity group '%s' when creating vm",
			name)
	}
	return affiId, nil
}

func (a CPI) generateNetworksMap(networks apiv1.Networks, defNetName, zoneId string) ([]map[string]string, error) {
	netList := make([]map[string]string, 0)
	for _, network := range networks {
		if network.Type() == string(config.VipNetwork) {
			continue
		}
		var networkProps NetworkCloudProperties
		err := network.CloudProps().As(&networkProps)
		if err != nil {
			return netList, err
		}
		if networkProps.Name == defNetName {
			continue
		}
		netCs, err := a.findNetworkByName(networkProps.Name, zoneId)
		if err != nil {
			return netList, bosherr.WrapErrorf(err, "Could not found network %s when creating vm", networkProps.Name)
		}
		netList = append(netList, a.networkToMap(netCs, network.Type(), network.IP()))
	}
	return netList, nil
}

func (a CPI) networkToMap(net *cloudstack.Network, netType, ip string) map[string]string {
	m := make(map[string]string)
	m["networkid"] = net.Id
	if netType != string(config.DynamicNetwork) {
		m["ip"] = ip
	}
	return m
}
