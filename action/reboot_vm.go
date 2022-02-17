package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) RebootVM(cid apiv1.VMCID) error {
	a.logger.Info("reboot_vm", "rebooting vm '%s' ...", cid.AsString())
	err := a.rebootVM(cid.AsString())
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not reboot vm '%s'", cid.AsString())
		a.logger.Error("reboot_vm", err.Error())
		return err
	}
	a.logger.Info("reboot_vm", "finished rebooting vm '%s'", cid.AsString())
	return nil
}

func (a CPI) rebootVM(vmName string) error {
	vm, err := a.vmFindByName(vmName)
	if err != nil {
		return err
	}
	return a.vmReboot(vm)
}
