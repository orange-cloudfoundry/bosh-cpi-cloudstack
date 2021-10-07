package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SetVMMetadata(cid apiv1.VMCID, meta apiv1.VMMeta) error {
	a.logger.Info("set_vm_metadata", "Creating vm metadata for %s", cid)
	vm, err := a.findVmByName(cid)
	if err != nil {
		bosherr.WrapErrorf(err, "Setting metadata for vm")
	}
	err = a.setMetadata(config.UserVm, vm.Id, &meta)
	a.logger.Info("set_vm_metadata", "Finished creating vm metadata for %s", cid)
	return err
}
