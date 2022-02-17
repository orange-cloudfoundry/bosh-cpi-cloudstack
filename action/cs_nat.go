package action

import (
	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)


func (a CPI) natIFRFindByVM(vm *cloudstack.VirtualMachine) ([]*cloudstack.IpForwardingRule, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.Global)

	a.logger.Debug("natIFRFindByVM", "listing ip forward rules for vm '%s' (%s)...", vm.Name, vm.Id)
	p := a.client.NAT.NewListIpForwardingRulesParams()
	p.SetVirtualmachineid(vm.Id)
	resp, err := a.client.NAT.ListIpForwardingRules(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not free virtual ip rules for vm %s (%s)", vm.Name, vm.Id)
		a.logger.Error("natIFRFindByVM", err.Error())
		return nil, err
	}

	a.logger.Debug("natIFRFindByVM", "finished listing ip forward rules for vm '%s' (%s)", vm.Name, vm.Id)
	return resp.IpForwardingRules, nil
}

func (a CPI) natIFRDelete(ifr *cloudstack.IpForwardingRule) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.Global)

	a.logger.Debug("natIFRDelete", "deleting ip forward rules '%s'...", ifr.Id)
	p := a.client.NAT.NewDeleteIpForwardingRuleParams(ifr.Id)
	err := a.client.NAT.DeleteIpForwardingRule(p)

	if err != nil {
		err = bosherr.WrapErrorf(err, "could not delete ip forwardng rule '%s'", ifr.Id)
		a.logger.Error("natIFRDelete", err.Error())
		return err
	}

	a.logger.Debug("natIFRDelete", "finished deleting ip forward rules '%s'", ifr.Id)
	return nil
}

func (a CPI) natIFRsDelete(ifrs []*cloudstack.IpForwardingRule) error {
	for _, cIFR := range ifrs {
		if err := a.natIFRDelete(cIFR); err != nil {
			return err
		}
	}
	return nil
}


func (a CPI) natDisableForIp(ip *cloudstack.PublicIpAddress) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.Global)

	a.logger.Debug("natDisableForIp", "disabling static nat for ip '%s'...", ip.Id)
	p := a.client.NAT.NewDisableStaticNatParams(ip.Id)
	err := a.client.NAT.DisableStaticNat(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not disable static nat for ip '%s'", ip.Id)
		a.logger.Error("natDisableForIp", err.Error())
		return err
	}

	a.logger.Debug("natDisableForIp", "finished disabling static nat for ip '%s'", ip.Id)
	return nil
}


func (a CPI) natDisableForIps(ips []*cloudstack.PublicIpAddress) error {
	for _, cIp := range ips {
		if err := a.natDisableForIp(cIp); err != nil {
			return err
		}
	}
	return nil
}
