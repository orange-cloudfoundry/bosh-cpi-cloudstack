package action

func (a CPI) setLoadBalancers(lbProps LBCloudProperties, vmID, zoneId, networkId string) error {
	for _, lbConf := range lbProps.Lbs {
		err := a.setLoadBalancer(lbConf, vmID, zoneId, networkId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a CPI) setLoadBalancer(lbConf LBConfig, vmID, zoneId, networkId string) error {
	lbId := ""
	lb, err := a.findLBRuleByName(lbConf.Name)
	if err != nil {
		return err
	}
	if lb != nil {
		lbId = lb.Id
	}
	if lbId == "" {
		lbId, err = a.createLoadBalancer(lbConf, zoneId, networkId)
		if err != nil {
			return err
		}
	}
	p := a.client.LoadBalancer.NewAssignToLoadBalancerRuleParams(lbId)
	p.SetVirtualmachineids([]string{vmID})
	_, err = a.client.LoadBalancer.AssignToLoadBalancerRule(p)
	return err
}

func (a CPI) createLoadBalancer(lbConf LBConfig, zoneId, networkId string) (string, error) {
	p := a.client.LoadBalancer.NewCreateLoadBalancerRuleParams(lbConf.Algorithm, lbConf.Name, lbConf.PrivatePort, lbConf.PublicPort)
	p.SetZoneid(zoneId)
	p.SetOpenfirewall(lbConf.OpenFirewall)
	if lbConf.PublicIp != "" {
		publicIp, err := a.findPublicIpByIp(lbConf.PublicIp)
		if err != nil {
			return "", err
		}
		p.SetPublicipid(publicIp.Id)
	} else {
		p.SetNetworkid(networkId)
	}
	resp, err := a.client.LoadBalancer.CreateLoadBalancerRule(p)
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
