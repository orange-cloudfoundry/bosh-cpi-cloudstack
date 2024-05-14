package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SetVMMetadata(cid apiv1.VMCID, meta apiv1.VMMeta) error {
	a.logger.Info("set_vm_metadata", "Creating vm metadata for %s", cid)
	vm, err := a.findVmByName(cid)
	if err != nil {
		if err2 := bosherr.WrapErrorf(err, "Setting metadata for vm"); err != nil {
			a.logger.Error("set_vm_metadata", "double error while setting VM[%s] metadata: %s (%s)", cid, err2, err)
		}
	}
	err = a.setMetadata(config.UserVm, vm.Id, &meta)
	a.logger.Info("set_vm_metadata", "Finished creating vm metadata for %s", cid)
	return err
}
