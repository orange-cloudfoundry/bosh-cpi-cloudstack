package action

import (
	"fmt"
	"strings"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func (a CPI) setMetadata(tagType config.Tags, cid string, meta util.MetaMarshal) error {
	// listParams := a.client.Resourcetags.NewListTagsParams()
	// listParams.SetResourcetype(string(tagType))
	// listParams.SetResourceid(cid)

	// resp, _ := a.client.Resourcetags.ListTags(listParams)
	// if resp != nil && len(resp.Tags) > 0 {
	// 	_, err := a.client.Resourcetags.DeleteTags(a.client.Resourcetags.NewDeleteTagsParams([]string{cid}, string(tagType)))
	// 	if err != nil {
	// 		return bosherr.WrapErrorf(err, "Updating %s metadata '%s'", tagType, cid)
	// 	}
	// }
	// tags := util.ConvertMapToTags(meta)
	// tags["director_uuid"] = a.ctx.DirectorUUID
	params := a.client.Resourcetags.NewCreateTagsParams([]string{cid}, string(tagType), tags)
	_, err := a.client.Resourcetags.CreateTags(params)
	if err != nil {
		return bosherr.WrapErrorf(err, "Setting %s metadata '%s'", tagType, cid)
	}
	return nil
}


// func (a CPI) findServiceOfferingByName(name string) (*cloudstack.ServiceOffering, error) {
// 	p := a.client.ServiceOffering.NewListServiceOfferingsParams()
// 	p.SetName(name)

// 	resp, err := a.client.ServiceOffering.ListServiceOfferings(p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.ServiceOfferings) == 0 {
// 		return nil, fmt.Errorf("service offering '%s' not found", name)
// 	}
// 	return resp.ServiceOfferings[0], nil
// }

// func (a CPI) findNetworkOfferingByNetwork(networkId, zoneID string) (*cloudstack.NetworkOffering, error) {
// 	p := a.client.NetworkOffering.NewListNetworkOfferingsParams()
// 	p.SetZoneid(zoneID)
// 	p.SetNetworkid(networkId)

// 	resp, err := a.client.NetworkOffering.ListNetworkOfferings(p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.NetworkOfferings) == 0 {
// 		return nil, fmt.Errorf("Cannot found network offering by network id %s", networkId)
// 	}
// 	return resp.NetworkOfferings[0], nil
// }

// func (a CPI) findNetworkByName(name string, zoneID string) (*cloudstack.Network, error) {
// 	p := a.client.Network.NewListNetworksParams()
// 	p.SetZoneid(zoneID)

// 	resp, err := a.client.Network.ListNetworks(p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.Networks) == 0 {
// 		return nil, fmt.Errorf("Cannot found network %s", name)
// 	}

// 	for _, network := range resp.Networks {
// 		if network.Name == name {
// 			return network, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("Cannot found network %s", name)
// }

// func (a CPI) findTemplateByName(name string) (*cloudstack.Template, error) {
// 	p := a.client.Template.NewListTemplatesParams("executable")
// 	p.SetName(name)

// 	resp, err := a.client.Template.ListTemplates(p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.Templates) == 0 {
// 		return nil, fmt.Errorf("template '%s' not found", name)
// 	}
// 	return resp.Templates[0], nil
// }

// func (a CPI) findLBRuleByName(name string) (*cloudstack.LoadBalancerRule, error) {
// 	p := a.client.LoadBalancer.NewListLoadBalancerRulesParams()
// 	p.SetName(name)

// 	resp, err := a.client.LoadBalancer.ListLoadBalancerRules(p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.LoadBalancerRules) == 0 {
// 		return nil, nil
// 	}
// 	return resp.LoadBalancerRules[0], nil
// }

// func (a CPI) findPublicIpByIp(ip string) (*cloudstack.PublicIpAddress, error) {
// 	p := a.client.Address.NewListPublicIpAddressesParams()
// 	p.SetIpaddress(ip)
// 	resp, err := a.client.Address.ListPublicIpAddresses(p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.PublicIpAddresses) == 0 {
// 		return nil, fmt.Errorf("Cannot found public ip %s", ip)
// 	}
// 	return resp.PublicIpAddresses[0], nil
// }

// func (a CPI) findAffinityGroup(name string) (*cloudstack.AffinityGroup, error) {
// 	lsP := a.client.AffinityGroup.NewListAffinityGroupsParams()
// 	lsP.SetName(name)
// 	lsResp, err := a.client.AffinityGroup.ListAffinityGroups(lsP)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(lsResp.AffinityGroups) > 0 {
// 		return lsResp.AffinityGroups[0], nil
// 	}
// 	return nil, nil
// }

// func (a CPI) findOrCreateAffinityGroup(name, aType string) (string, error) {
// 	af, err := a.findAffinityGroup(name)
// 	if err != nil {
// 		return "", err
// 	}
// 	if af != nil && af.Type == aType {
// 		return af.Id, nil
// 	}

// 	// if user decide to change the affinity group type, we delete before redoing
// 	if af != nil && af.Type != aType {
// 		p := a.client.AffinityGroup.NewDeleteAffinityGroupParams()
// 		p.SetName(name)
// 		a.client.AffinityGroup.DeleteAffinityGroup(p)
// 	}
// 	p := a.client.AffinityGroup.NewCreateAffinityGroupParams(name, aType)
// 	resp, err := a.client.AffinityGroup.CreateAffinityGroup(p)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "already exists") {
// 			af, err := a.findAffinityGroup(name)
// 			id := ""
// 			if af != nil {
// 				id = af.Id
// 			}
// 			return id, err
// 		}
// 		return "", err
// 	}
// 	return resp.Id, nil
// }

// func retryable(timeout time.Duration, cmd func() error, checkRetry func(err error) bool) error {
// 	var timer time.Duration
// 	var err error
// 	currentTime := time.Now().Unix()
// 	for {
// 		err = cmd()
// 		if err == nil {
// 			return nil
// 		}
// 		if !checkRetry(err) {
// 			return err
// 		}
// 		if time.Now().Unix()-currentTime > int64(timeout) {
// 			return err
// 		}
// 		if timer < 15 {
// 			timer++
// 		}
// 		time.Sleep(timer)
// 	}
// 	return err
// }
