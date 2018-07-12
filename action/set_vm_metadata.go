package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SetVMMetadata(cid apiv1.VMCID, meta apiv1.VMMeta) error {
	return a.setMetadata(config.UserVm, cid.AsString(), &meta)
}
