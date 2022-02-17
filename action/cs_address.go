package action

import (
	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)


func (a CPI) addressListPublicIPs(allocatedOnly bool) ([]*cloudstack.PublicIpAddress, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.Global)

	a.logger.Debug("addressListPublicIPs", "listing public ip adresses...")
	p := a.client.Address.NewListPublicIpAddressesParams()
	p.SetAllocatedonly(allocatedOnly)
	resp, err := a.client.Address.ListPublicIpAddresses(p)

	if err != nil {
		err = bosherr.WrapErrorf(err, "could not list public adresses")
		a.logger.Error("addressListPublicIPs", err.Error())
		return err
	}

	a.logger.Debug("addressListPublicIPs", "finished listing public ip adresses")
	return resp.PublicIpAddresses, nil
}

func (a CPI) addressListPublicIPsByVM(vm *cloudstack.VirtualMachine) ([]*cloudstack.PublicIpAddress, error) {
	ips, err := a.addressListPublicIPs(true)
	if err != nil {
		return nil, err
	}
	res := []*cloudstack.PublicIpAddress{}
	for _, cIp := range ips {
		if cIp.Virtualmachineid == vm.Id {
			res = append(res, cIp)
		}
	}
	return res, nil
}
