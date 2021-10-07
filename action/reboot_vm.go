package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) RebootVM(cid apiv1.VMCID) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.RebootVm)

	vm, err := a.findVmByName(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Rebooting vm %s", cid.AsString())
	}

	a.logger.Info("reboot_vm", "Rebooting vm %s ...", cid.AsString())
	p := a.client.VirtualMachine.NewRebootVirtualMachineParams(vm.Id)
	_, err = a.client.VirtualMachine.RebootVirtualMachine(p)
	if err != nil {
		return bosherr.WrapErrorf(err, "Rebooting vm %s", cid.AsString())
	}
	a.logger.Info("reboot_vm", "Finished rebooting vm %s .", cid.AsString())
	return nil
}
