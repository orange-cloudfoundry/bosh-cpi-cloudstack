package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SetVMMetadata(cid apiv1.VMCID, meta apiv1.VMMeta) error {
	a.logger.Info("set_vm_metadata", "setting metadata on vm '%s'...", cid.AsString())
	err := a.setVMMetadata(cid.AsString(), &meta)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not set metadata on vm '%s'", cid.AsString())
		a.logger.Error("set_vm_metadata", err.Error())
		return err
	}
	a.logger.Info("set_vm_metadata", "finished setting metadata on vm '%s'...", cid.AsString())
	return nil
}


// setVMMetadata -
// 1. delete any tag that are already present
func (a CPI) setVMMetadata(vmName, meta util.MetaMarshal) error {
	vm, err := a.vmFindByName(vmName)
	if err != nil {
		return err
	}
	return a.setMetadata(string(config.UserVm), vm.Id, meta)
}
