package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"fmt"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
	"encoding/json"
	"encoding/base64"
	"github.com/xanzy/go-cloudstack/cloudstack"
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
	deplParams.SetNetworkids([]string{network.Id})
	deplParams.SetKeypair(a.config.CloudStack.DefaultKeyName)

	if !defaultNetwork.IsDynamic() {
		deplParams.SetIpaddress(defaultNetwork.IP())
	}

	affinId, err := a.generateAffinityGroup(resProps, env)
	if err != nil {
		return apiv1.VMCID{}, err
	}
	if affinId != "" {
		deplParams.SetAffinitygroupids([]string{affinId})
	}

	if a.config.CloudStack.RootDiskSize > 0 {
		deplParams.SetRootdisksize(a.config.CloudStack.RootDiskSize / 1024)
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
	diskCid, err := a.createEphemeralDisk(resProps.EphemeralDiskSize, resProps.DiskCloudProperties, nil)
	if err != nil {
		return apiv1.VMCID{}, a.destroyVmErrFallback(bosherr.WrapError(err, "Cannot create ephemeral disk when creating vm"), resp.Id)
	}
	a.logger.Info("create_vm", "Finished creating ephemeral disk for vm %s .", vmName)

	a.logger.Info("create_vm", "Attaching ephemeral disk for vm %s ...", vmName)
	err = a.AttachDisk(vmCID, diskCid)
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

	_, err = a.client.NAT.EnableStaticNat(p)
	return err
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
	if (nbDynamic + nbManual) > 1 {
		return bosherr.Errorf("Only 1 nic is supported, mixing manual and dynamic network is not allowed")
	}
	if (nbDynamic + nbManual) == 0 {
		return bosherr.Errorf("It must have one dynamic or one manual network defined")
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
