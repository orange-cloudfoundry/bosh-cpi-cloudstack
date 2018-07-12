package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) SetDiskMetadata(cid apiv1.DiskCID, meta apiv1.DiskMeta) error {
	vol, err := a.findVolumeById(cid)
	if err != nil {
		bosherr.WrapErrorf(err, "Setting metadata for volume")
	}
	return a.setMetadata(config.Volume, vol.Id, &meta)
}
