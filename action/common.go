package action

import (
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
	"github.com/xanzy/go-cloudstack/cloudstack"
	"strings"
	"time"
)

func (a CPI) setMetadata(tagType config.Tags, cid string, meta util.MetaMarshal) error {
	listParams := a.client.Resourcetags.NewListTagsParams()
	listParams.SetResourcetype(string(tagType))
	listParams.SetResourceid(cid)

	resp, _ := a.client.Resourcetags.ListTags(listParams)
	if resp != nil && len(resp.Tags) > 0 {
		_, err := a.client.Resourcetags.DeleteTags(a.client.Resourcetags.NewDeleteTagsParams([]string{cid}, string(tagType)))
		if err != nil {
			return bosherr.WrapErrorf(err, "Updating %s metadata '%s'", tagType, cid)
		}
	}
	tags := util.ConvertMapToTags(meta)
	tags["director_uuid"] = a.ctx.DirectorUUID
	params := a.client.Resourcetags.NewCreateTagsParams([]string{cid}, string(tagType), tags)
	_, err := a.client.Resourcetags.CreateTags(params)
	if err != nil {
		return bosherr.WrapErrorf(err, "Setting %s metadata '%s'", tagType, cid)
	}
	return nil
}

func (a CPI) findVmsByName(cid apiv1.VMCID) ([]*cloudstack.VirtualMachine, error) {
	p := a.client.VirtualMachine.NewListVirtualMachinesParams()
	p.SetName(cid.AsString())
	resp, err := a.client.VirtualMachine.ListVirtualMachines(p)
	if err != nil {
		return []*cloudstack.VirtualMachine{}, err
	}
	return resp.VirtualMachines, nil
}

func (a CPI) findVolumesByName(cid apiv1.DiskCID) ([]*cloudstack.Volume, error) {
	p := a.client.Volume.NewListVolumesParams()
	p.SetName(cid.AsString())
	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		return []*cloudstack.Volume{}, err
	}
	return resp.Volumes, nil
}

func (a CPI) findVmByName(cid apiv1.VMCID) (*cloudstack.VirtualMachine, error) {
	vms, err := a.findVmsByName(cid)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Can't find vm name '%s'", cid.AsString())
	}
	if len(vms) == 0 {
		return nil, bosherr.Errorf("Can't find vm name '%s'", cid.AsString())

	}
	return vms[0], nil
}

func (a CPI) findVolumeByName(cid apiv1.DiskCID) (*cloudstack.Volume, error) {
	volumes, err := a.findVolumesByName(cid)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Can't find disk name '%s'", cid.AsString())
	}
	if len(volumes) == 0 {
		return nil, bosherr.Errorf("Can't find disk name '%s'", cid.AsString())
	}
	return volumes[0], nil
}

func (a CPI) findZoneID() (string, error) {
	p := a.client.Zone.NewListZonesParams()
	p.SetName(a.config.CloudStack.DefaultZone)
	resp, err := a.client.Zone.ListZones(p)
	if err != nil {
		return "", err
	}
	if len(resp.Zones) == 0 {
		return "", bosherr.Errorf("zone '%s' not found", a.config.CloudStack.DefaultZone)
	}
	return resp.Zones[0].Id, nil
}

func (a CPI) findDiskOfferingByName(name string) (*cloudstack.DiskOffering, error) {
	p := a.client.DiskOffering.NewListDiskOfferingsParams()
	p.SetName(name)

	resp, err := a.client.DiskOffering.ListDiskOfferings(p)
	if err != nil {
		return nil, err
	}
	if len(resp.DiskOfferings) == 0 {
		return nil, fmt.Errorf("Cannot found disk offering %s", name)
	}
	return resp.DiskOfferings[0], nil
}

func (a CPI) findServiceOfferingByName(name string) (*cloudstack.ServiceOffering, error) {
	p := a.client.ServiceOffering.NewListServiceOfferingsParams()
	p.SetName(name)

	resp, err := a.client.ServiceOffering.ListServiceOfferings(p)
	if err != nil {
		return nil, err
	}
	if len(resp.ServiceOfferings) == 0 {
		return nil, fmt.Errorf("service offering '%s' not found", name)
	}
	return resp.ServiceOfferings[0], nil
}

func (a CPI) findNetworkOfferingByNetwork(networkId, zoneID string) (*cloudstack.NetworkOffering, error) {
	p := a.client.NetworkOffering.NewListNetworkOfferingsParams()
	p.SetZoneid(zoneID)
	p.SetNetworkid(networkId)

	resp, err := a.client.NetworkOffering.ListNetworkOfferings(p)
	if err != nil {
		return nil, err
	}
	if len(resp.NetworkOfferings) == 0 {
		return nil, fmt.Errorf("Cannot found network offering by network id %s", networkId)
	}
	return resp.NetworkOfferings[0], nil
}

func (a CPI) findNetworkByName(name string, zoneID string) (*cloudstack.Network, error) {
	p := a.client.Network.NewListNetworksParams()
	p.SetZoneid(zoneID)

	resp, err := a.client.Network.ListNetworks(p)
	if err != nil {
		return nil, err
	}
	if len(resp.Networks) == 0 {
		return nil, fmt.Errorf("Cannot found network %s", name)
	}

	for _, network := range resp.Networks {
		if network.Name == name {
			return network, nil
		}
	}
	return nil, fmt.Errorf("Cannot found network %s", name)
}

func (a CPI) findTemplateByName(name string) (*cloudstack.Template, error) {
	p := a.client.Template.NewListTemplatesParams("executable")
	p.SetName(name)

	resp, err := a.client.Template.ListTemplates(p)
	if err != nil {
		return nil, err
	}
	if len(resp.Templates) == 0 {
		return nil, fmt.Errorf("template '%s' not found", name)
	}
	return resp.Templates[0], nil
}

func (a CPI) findLBRuleByName(name string) (*cloudstack.LoadBalancerRule, error) {
	p := a.client.LoadBalancer.NewListLoadBalancerRulesParams()
	p.SetName(name)

	resp, err := a.client.LoadBalancer.ListLoadBalancerRules(p)
	if err != nil {
		return nil, err
	}
	if len(resp.LoadBalancerRules) == 0 {
		return nil, nil
	}
	return resp.LoadBalancerRules[0], nil
}

func (a CPI) findOsTypeId(descr string) (string, error) {
	p := a.client.GuestOS.NewListOsTypesParams()
	p.SetDescription(descr)
	resp, err := a.client.GuestOS.ListOsTypes(p)
	if err != nil {
		return "", bosherr.WrapErrorf(err, "Unable to list guest os types")
	}
	if resp.Count == 0 {
		return "", bosherr.WrapErrorf(err, "Can't find guest os type '%s'", descr)
	}
	return resp.OsTypes[0].Id, nil
}

func (a CPI) findPublicIpByIp(ip string) (*cloudstack.PublicIpAddress, error) {
	p := a.client.Address.NewListPublicIpAddressesParams()
	p.SetIpaddress(ip)
	resp, err := a.client.Address.ListPublicIpAddresses(p)
	if err != nil {
		return nil, err
	}
	if len(resp.PublicIpAddresses) == 0 {
		return nil, fmt.Errorf("Cannot found public ip %s", ip)
	}
	return resp.PublicIpAddresses[0], nil
}

func (a CPI) findAffinityGroup(name, aType string) (string, error) {
	lsP := a.client.AffinityGroup.NewListAffinityGroupsParams()
	lsP.SetName(name)
	lsResp, err := a.client.AffinityGroup.ListAffinityGroups(lsP)
	if err != nil {
		return "", err
	}
	if len(lsResp.AffinityGroups) > 0 {
		return lsResp.AffinityGroups[0].Id, nil
	}
	return "", nil
}

func (a CPI) findOrCreateAffinityGroup(name, aType string) (string, error) {
	afId, err := a.findAffinityGroup(name, aType)
	if err != nil {
		return "", err
	}
	if afId != "" {
		return afId, nil
	}

	p := a.client.AffinityGroup.NewCreateAffinityGroupParams(name, aType)
	resp, err := a.client.AffinityGroup.CreateAffinityGroup(p)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return a.findAffinityGroup(name, aType)
		}
		return "", err
	}
	return resp.Id, nil
}

func retryable(timeout time.Duration, cmd func() error, checkRetry func(err error) bool) error {
	var timer time.Duration
	var err error
	currentTime := time.Now().Unix()
	for {
		err = cmd()
		if err == nil {
			return nil
		}
		if !checkRetry(err) {
			return err
		}
		if time.Now().Unix()-currentTime > int64(timeout) {
			return err
		}
		if timer < 15 {
			timer++
		}
		time.Sleep(timer)
	}
	return err
}
