package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) SetVMMetadata(cid apiv1.VMCID, meta apiv1.VMMeta) error {
	vm, err := a.findVmByName(cid)
	if err != nil {
		bosherr.WrapErrorf(err, "Setting metadata for vm")
	}
	return a.setMetadata(config.UserVm, vm.Id, &meta)
}
