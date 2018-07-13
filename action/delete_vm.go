package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) DeleteVM(cid apiv1.VMCID) error {
	vms, err := a.findVmsByName(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when finding vm %s", cid.AsString())
	}

	if len(vms) > 1 {
		return bosherr.Errorf("Too much vms with name %s", cid.AsString())
	}

	if len(vms) == 0 {
		return bosherr.Errorf("No vm found with name %s", cid.AsString())
	}
	vm := vms[0]

	err = a.liberateVIPs(vm.Id)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when liberating vips for vm %s", cid.AsString())
	}

	err = a.stopVmById(vm.Id)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when stopping vm %s", cid.AsString())
	}

	err = a.detachAllDisks(vm.Id)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when detaching all volumes for vm %s", cid.AsString())
	}

	err = a.deleteVMById(vm.Id)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when destroying vm %s", cid.AsString())
	}

	a.regFactory.Create(cid).Delete()

	return nil
}

func (a CPI) stopVmById(vmId string) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.StopVm)

	p := a.client.VirtualMachine.NewStopVirtualMachineParams(vmId)
	_, err := a.client.VirtualMachine.StopVirtualMachine(p)
	return err
}

func (a CPI) deleteVMById(vmId string) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteVm)

	p := a.client.VirtualMachine.NewDestroyVirtualMachineParams(vmId)
	_, err := a.client.VirtualMachine.DestroyVirtualMachine(p)
	return err
}

func (a CPI) detachAllDisks(vmId string) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DetachVolume)
	listParams := a.client.Volume.NewListVolumesParams()
	listParams.SetVirtualmachineid(vmId)
	listParams.SetType(string(config.Datadisk))

	resp, err := a.client.Volume.ListVolumes(listParams)
	if err != nil {
		return bosherr.WrapErrorf(err, "Cannot get volumes for vm id %s", vmId)
	}

	for _, volume := range resp.Volumes {
		detachParams := a.client.Volume.NewDetachVolumeParams()
		detachParams.SetId(volume.Id)
		_, err = a.client.Volume.DetachVolume(detachParams)
		if err != nil {
			return bosherr.WrapErrorf(err, "Cannot detach volume %s for vm id %s", volume.Name, vmId)
		}
	}
	return nil
}

func (a CPI) liberateVIPs(vmId string) error {
	ruleParams := a.client.NAT.NewListIpForwardingRulesParams()
	ruleParams.SetVirtualmachineid(vmId)
	respRules, err := a.client.NAT.ListIpForwardingRules(ruleParams)
	if err != nil {
		return bosherr.WrapErrorf(err, "Can't liberate vip rules for vm id %s", vmId)
	}
	for _, rule := range respRules.IpForwardingRules {
		ruleDelParams := a.client.NAT.NewDeleteIpForwardingRuleParams(rule.Id)
		a.client.NAT.DeleteIpForwardingRule(ruleDelParams)
	}

	listPubIpParams := a.client.Address.NewListPublicIpAddressesParams()
	listPubIpParams.SetAllocatedonly(true)
	respPubIp, err := a.client.Address.ListPublicIpAddresses(listPubIpParams)
	if err != nil {
		return bosherr.WrapErrorf(err, "Can't liberate vip for vm id %s", vmId)
	}

	for _, pubIp := range respPubIp.PublicIpAddresses {
		if pubIp.Virtualmachineid != vmId {
			continue
		}
		disableParams := a.client.NAT.NewDisableStaticNatParams(pubIp.Id)
		a.client.NAT.DisableStaticNat(disableParams)
	}
	return nil
}
