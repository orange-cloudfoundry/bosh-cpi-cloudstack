package action

import (
	"time"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
)


// setMetadata - associate tags to given object
// 1. delete any tag that are already present
func (a CPI) setMetadata(objectType string, objectID string, meta util.MetaMarshal) error {
	// 1.
	tags, err := a.tagList(objectType, objectID)
	if (err == nil) && (len(tags) != 0) {
		err = a.tagDelete(objectType, objectID)
		if err != nil {
			return bosherr.WrapErrorf(err, "could not delete pre-existing tags for %s:%s", objectType, objectID)
		}
	}

	tagMap := util.ConvertMapToTags(meta)
	tagMap["director_uuid"] = a.ctx.DirectorUUID
	return a.tagCreate(objectType, objectID, tagMap)
}


func (a CPI) retryable(timeout time.Duration, cmd func() error, checkRetry func(err error) bool) error {
	var (
		timer time.Duration
		err error
	)

	count := 0
	currentTime := time.Now().Unix()

	a.logger.Debug("retryable", "starting retryable call...")
	for {
		a.logger.Debug("retryable", "running retry number '%d'", count)
		err = cmd()
		if err == nil {
			a.logger.Debug("retryable", "finished with successful operation")
			return nil
		}
		if !checkRetry(err) {
			a.logger.Debug("retryable", "finished with error, operation is not retryable")
			return err
		}
		if time.Now().Unix()-currentTime > int64(timeout) {
			a.logger.Debug("retryable", "finished with error, operation timeout")
			return err
		}
		if timer < 15 {
			timer++
		}
		time.Sleep(timer)
		count++
	}
	return err
}


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
