package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SetDiskMetadata(cid apiv1.DiskCID, meta apiv1.DiskMeta) error {
	vol, err := a.findVolumeByName(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Setting metadata for volume")
	}
	return a.setMetadata(config.Volume, vol.Id, &meta)
}
