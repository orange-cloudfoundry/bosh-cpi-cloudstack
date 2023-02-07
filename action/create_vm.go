package action

import (
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	uuid "github.com/satori/go.uuid"
)

const (
	pvDriverErr         = "VM which requires PV drivers to be installed"
	defaultAffinityType = "host anti-affinity"
)

type CreateArgs struct {
	agentID            apiv1.AgentID
	stemcellCID        apiv1.StemcellCID
	cloudProps         apiv1.VMCloudProps
	networks           apiv1.Networks
	associatedDiskCIDs []apiv1.DiskCID
	env                apiv1.VMEnv
}

func (a CPI) CreateVM(
	agentID apiv1.AgentID,
	stemcellCID apiv1.StemcellCID,
	cloudProps apiv1.VMCloudProps,
	networks apiv1.Networks,
	associatedDiskCIDs []apiv1.DiskCID,
	env apiv1.VMEnv,
) (apiv1.VMCID, error) {

	args := CreateArgs{agentID, stemcellCID, cloudProps, networks, associatedDiskCIDs, env}
	vmCID, _, err := a.CreateBase(args, false)
	return vmCID, err
}

func (a CPI) CreateVMV2(
	agentID apiv1.AgentID,
	stemcellCID apiv1.StemcellCID,
	cloudProps apiv1.VMCloudProps,
	networks apiv1.Networks,
	associatedDiskCIDs []apiv1.DiskCID,
	env apiv1.VMEnv) (apiv1.VMCID, apiv1.Networks, error) {

	args := CreateArgs{agentID, stemcellCID, cloudProps, networks, associatedDiskCIDs, env}
	return a.CreateBase(args, true)
}

func (a CPI) CreateBase(p CreateArgs, isV2 bool) (apiv1.VMCID, apiv1.Networks, error) {
	var resProps ResourceCloudProperties

	a.client.AsyncTimeout(a.config.CloudStack.Timeout.CreateVm)
	a.client.Timeout(time.Duration(a.config.CloudStack.Timeout.CreateVm) * time.Second)

	vmName := fmt.Sprintf("%s%s", config.VMPrefix, uuid.NewV4().String())
	vmCID := apiv1.NewVMCID(vmName)

	if err := p.cloudProps.As(&resProps); err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapError(err, "Cannot create vm")
	}

	zoneID, err := a.findZoneID()
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapError(err, "could not found zone")
	}

	serviceOffering, err := a.findServiceOfferingByName(resProps.ComputeOffering)
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapError(err, "could not found compute offering")
	}

	template, err := a.findTemplateByName(p.stemcellCID.AsString())
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapError(err, "could not found template")
	}

	if err = a.checkNetworkConfig(p.networks); err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapErrorf(err, "invalid network configuration")
	}

	if err = a.addRouteToNetworks(resProps.Routes, p.networks); err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapErrorf(err, "invalid network configuration")
	}

	// computing create vm paramters
	a.logger.Info("create_vm", "Computing vm deploy parameters %s ...", vmName)
	deplParams := a.client.VirtualMachine.NewDeployVirtualMachineParams(serviceOffering.Id, template.Id, zoneID)
	deplParams.SetName(vmName)
	deplParams.SetKeypair(a.config.CloudStack.DefaultKeyName)

	// [params] create
	if serviceOffering.Iscustomized {
		cpu := resProps.CPUNumber
		cpuSpeed := resProps.CPUSpeed
		ram := resProps.RAM
		if 0 == cpuSpeed {
			cpuSpeed = 2000
		}
		if (0 == cpu) || (0 == ram) {
			return apiv1.VMCID{}, apiv1.Networks{}, bosherr.Errorf("Could not find `cpu` and `memory` cloud_properties mandatory for compute offering %s when creating vm", resProps.ComputeOffering)
		}
		deplParams.SetDetails(map[string]string{
			"cpuNumber": fmt.Sprintf("%d", cpu),
			"cpuSpeed":  fmt.Sprintf("%d", cpuSpeed),
			"memory":    fmt.Sprintf("%d", ram),
		})
	}

	// [params] create networks
	networkDefault, networkMaps, err := a.generateNetworksMap(p.networks, zoneID)
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapError(err, "unable to generate network configuration")
	}
	for _, network := range networkMaps {
		deplParams.AddIptonetworklist(network)
	}

	// [params] create user-data
	userDataService := NewUserDataService(a.logger, vmName, a.config.Actions.Registry, p.networks, isV2)
	userDataService.SetAgentSettings(p.agentID, vmCID, p.networks, p.env, a.config.Actions.Agent)
	deplParams.SetUserdata(userDataService.ToBase64())

	affinID, err := a.generateAffinityGroup(resProps, p.env)
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, err
	}
	if affinID != "" {
		deplParams.SetAffinitygroupids([]string{affinID})
	}
	if resProps.RootDiskSize > 0 {
		deplParams.SetRootdisksize(int64(resProps.RootDiskSize / 1024))
	}
	a.logger.Debug("create_vm", "deploy parameters %#v", deplParams)
	a.logger.Info("create_vm", "Finished computing vm deploy parameters %s ...", vmName)

	// creating virtual machine
	a.logger.Info("create_vm", "Creating vm %s ...", vmName)
	resp, err := a.client.VirtualMachine.DeployVirtualMachine(deplParams)
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, bosherr.WrapError(err, "Error when creating vm")
	}
	a.logger.Debug("create_vm", "DeployVirtualMachine response: %#v", resp)
	if config.ToVmState(resp.State) != config.VmRunning {
		err = bosherr.Errorf("vm is not running after creation, actual state is %s", resp.State)
		return apiv1.VMCID{}, apiv1.Networks{}, a.destroyVmErrFallback(err, resp.Id)
	}
	a.logger.Info("create_vm", "Finished creating vm %s .", vmName)

	// process networks
	a.logger.Info("create_vm", "Post-processing networks...", vmName)
	if err = a.applyMacToNetworks(resp, p.networks); err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, a.destroyVmErrFallback(err, resp.Id)
	}
	a.logger.Info("create_vm", "Finished post-processing networks...", vmName)

	// populate registry
	if !isV2 {
		a.logger.Info("create_vm", "Registering vm %s in registry...", vmName)
		agentEnv := apiv1.NewAgentEnvFactory().ForVM(p.agentID, vmCID, p.networks, p.env, a.config.Actions.Agent)
		agentEnv.AttachSystemDisk(apiv1.NewDiskHintFromString("/dev/xvda"))
		agentEnv.AttachEphemeralDisk(apiv1.NewDiskHintFromString("/dev/xvdb"))
		envSvc := a.regFactory.Create(vmCID)
		val, _ := agentEnv.AsBytes()
		a.logger.Debug("create_vm", "Sending to registry %s", string(val))
		if err = envSvc.Update(agentEnv); err != nil {
			err := bosherr.WrapError(err, "unable to send data to registry")
			return apiv1.VMCID{}, apiv1.Networks{}, a.destroyVmErrFallback(err, resp.Id)
		}
		a.logger.Info("create_vm", "Finished registering vm %s in registry.", vmName)
	}

	// creating virtual IPs
	a.logger.Info("create_vm", "Creating vip(s) for vm %s ...", vmName)
	err = a.createVips(p.networks, resp.Id, zoneID, networkDefault)
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, a.destroyVmErrFallback(bosherr.WrapErrorf(err, "Could not create vips"), resp.Id)
	}
	a.logger.Info("create_vm", "Finished creating vip(s) for vm %s .", vmName)

	// creating ephemeral disks
	a.logger.Info("create_vm", "Creating ephemeral disk for vm %s ...", vmName)
	diskCid, err := a.createEphemeralDisk(resProps.EphemeralDiskSize, resProps.DiskCloudProperties, "")
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, a.destroyVmErrFallback(bosherr.WrapError(err, "Cannot create ephemeral disk when creating vm"), resp.Id)
	}
	a.logger.Info("create_vm", "Finished creating ephemeral disk for vm %s .", vmName)

	// creating load-balancers
	if len(resProps.Lbs) > 0 {
		a.logger.Info("create_vm", "Assigning vm %s to loadbalancers ...", vmName)
		err = a.setLoadBalancers(resProps.LBCloudProperties, resp.Id, zoneID, networkDefault.Id)
		if err != nil {
			return apiv1.VMCID{}, apiv1.Networks{}, a.destroyVmErrFallback(bosherr.WrapError(err, "Cannot assign loadbalancers to vm"), resp.Id)
		}
		a.logger.Info("create_vm", "Finished assigning vm %s to loadbalancers.", vmName)
	}

	// attaching disks
	a.logger.Info("create_vm", "Attaching ephemeral disk for vm %s ...", vmName)
	err = a.attachEphemeralDisk(vmCID, diskCid)
	if err != nil {
		return apiv1.VMCID{}, apiv1.Networks{}, a.destroyVmErrFallback(
			bosherr.WrapError(
				err,
				"Cannot attach ephemeral disk when creating vm"),
			resp.Id,
			func() {
				err := a.DeleteDisk(diskCid)
				if err != nil {
					a.logger.Warn("delete_disk", "unable to delete disk: %s", diskCid)
				}
			},
		)
	}
	a.logger.Info("create_vm", "Finished attaching ephemeral disk for vm %s .", vmName)

	return vmCID, p.networks, nil
}

func (a CPI) applyMacToNetworks(resp *cloudstack.DeployVirtualMachineResponse, boshNetworks apiv1.Networks) error {
	for _, nic := range resp.Nic {
		for name, network := range boshNetworks {
			if nic.Ipaddress == network.IP() {
				a.logger.Debug("create_vm", "setting mac address '%s' for network %s", nic.Macaddress, name)
				network.SetMAC(nic.Macaddress)
			}
		}
	}
	return nil
}

func (a CPI) parseCIDR(cidr string) (string, string, error) {
	addr, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", "", err
	}
	mask := fmt.Sprintf("%d.%d.%d.%d", ipNet.Mask[0], ipNet.Mask[1], ipNet.Mask[2], ipNet.Mask[3])
	return addr.String(), mask, nil
}

func (a CPI) addRouteToNetworks(routes VMExtRouteMap, boshNetworks apiv1.Networks) error {
	a.logger.Debug("create_vm", "routes from cloud config: %#v", routes)
	for networkName, routeList := range routes {
		targetNet, ok := boshNetworks[networkName]
		if !ok {
			return bosherr.Errorf("Invalid vm_extension: attempt to map route on unkown network '%s'", networkName)
		}
		for _, route := range routeList {
			ip, mask, err := a.parseCIDR(route.CIDR)
			if err != nil {
				return bosherr.Errorf("Invalid vm_extension: malformed CIDR '%s'", route)
			}
			if route.Gateway == "" {
				route.Gateway = targetNet.Gateway()
			}
			targetNet.AddRoute(ip, mask, route.Gateway)
		}
		boshNetworks[networkName] = targetNet
	}
	return nil
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

func (a CPI) createVips(networks apiv1.Networks, vmId, zoneID string, defNetwork *cloudstack.Network) error {
	for _, network := range networks {
		if network.Type() != string(config.VipNetwork) {
			continue
		}
		err := a.createVip(network, vmId, zoneID, defNetwork)
		if err != nil {
			return bosherr.WrapErrorf(err, "Error when creating vip %s", network.IP())
		}
	}
	return nil
}

func (a CPI) createVip(network apiv1.Network, vmId, zoneID string, defNetwork *cloudstack.Network) error {
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
		networkCs, err = a.findNetworkByName(networkProps.Name, zoneID)
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
		return bosherr.Errorf("invalid network configuration: only 1 vip is supported")
	}
	if (nbDynamic + nbManual) == 0 {
		return bosherr.Errorf("invalid network configuration: must have at least one dynamic or one manual network")
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
	affinityType := a.config.CloudStack.AutoAntiAffinityType
	if affinityType == "" {
		affinityType = defaultAffinityType
	}
	affiId, err := a.findOrCreateAffinityGroup(name, affinityType)
	if err != nil {
		return "", bosherr.WrapErrorf(
			err,
			"Could not find or create affinity group '%s' when creating vm",
			name)
	}
	return affiId, nil
}

type NetworkElem struct {
	Name       string
	Network    apiv1.Network
	Properties NetworkCloudProperties
}

// generateNetworksMap - create structure to populate CS create_vm parameters
//
// On CS stemcell, it is important that eth0 (ie: first interface) holds the
// metadata hdcp discovery
// We sort available network to make sure `cloud_properties: {"default":true}`
// networks are given as first interface to cloudstack create_vm parameters
//
// 1. create sortable structure
// 2. sort network, default first, the by name
// 3. create CS network list
//   - vip network not given to cloudstack
func (a CPI) generateNetworksMap(networks apiv1.Networks, zoneID string) (*cloudstack.Network, []map[string]string, error) {
	result := make([]map[string]string, 0)
	var defaultNetwork *cloudstack.Network

	// 1.
	data := make([]NetworkElem, 0)
	for name, network := range networks {
		var properties NetworkCloudProperties
		if err := network.CloudProps().As(&properties); err != nil {
			return nil, nil, bosherr.Errorf("invalid cloud_properties for network '%s'", name)
		}
		data = append(data, NetworkElem{
			Name:       name,
			Network:    network,
			Properties: properties,
		})
	}

	// 2.
	sort.SliceStable(data, func(i, j int) bool {
		usableFirst := (!data[i].Properties.UnDiscoverable) && (data[i].Network.Type() != string(config.VipNetwork))
		usableSecond := (!data[j].Properties.UnDiscoverable) && (data[j].Network.Type() != string(config.VipNetwork))
		if usableFirst && !usableSecond {
			return true
		} else if !usableFirst && usableSecond {
			return false
		}
		return data[i].Name < data[j].Name
	})

	// 3.
	index := 0
	for _, item := range data {
		if item.Network.Type() == string(config.VipNetwork) {
			continue
		}
		alias := fmt.Sprintf("eth%d", index)
		index += 1

		network := networks[item.Name]
		network.SetAlias(alias)
		networks[item.Name] = network

		csNet, err := a.findNetworkByName(item.Properties.Name, zoneID)
		if err != nil {
			return nil, nil, err
		}
		if defaultNetwork == nil {
			defaultNetwork = csNet
		}
		result = append(result, a.networkToMap(csNet, item.Network.Type(), item.Network.IP()))
	}
	return defaultNetwork, result, nil
}

func (a CPI) networkToMap(net *cloudstack.Network, netType, ip string) map[string]string {
	m := make(map[string]string)
	m["networkid"] = net.Id
	if netType != string(config.DynamicNetwork) {
		m["ip"] = ip
	}
	return m
}

// Local Variables:
// ispell-local-dictionary: "american"
// End:
