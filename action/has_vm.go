package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) HasVM(cid apiv1.VMCID) (bool, error) {
	a.logger.Info("has_vm", "checking vm exists for '%s'...", cid.AsString())
	found, err := a.hasVM(cid.AsString())
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not check if vm exists for '%s'", cid.AsString())
		a.logger.Error("has_vm", err.Error())
		return false, err
	}
	a.logger.Info("has_vm", "finished checking vm exists for '%s' (%t)", cid.AsString(), found)
	return found, nil
}


func (a CPI) hasVM(vmName string) (bool, error) {
	vms, err := a.vmsFindByName(vmName)
	if err != nil {
		return false, err
	}
	if len(vms) > 1 {
		return false, bosherr.Errorf("found multiple instances of mv '%s'", vmName)
	}
	return len(vms) == 1, nil
}
