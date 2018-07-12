package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SetDiskMetadata(cid apiv1.DiskCID, meta apiv1.DiskMeta) error {
	return a.setMetadata(config.Volume, cid.AsString(), &meta)
}
