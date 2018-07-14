package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"strings"
)

func (a CPI) SetDiskMetadata(cid apiv1.DiskCID, meta apiv1.DiskMeta) error {
	vol, err := a.findVolumeByName(cid)
	if err != nil {
		bosherr.WrapErrorf(err, "Setting metadata for volume")
	}
	a.setEphemeralDiskMetadata(vol.Virtualmachineid, meta)
	return a.setMetadata(config.Volume, vol.Id, &meta)
}

func (a CPI) setEphemeralDiskMetadata(vmId string, meta apiv1.DiskMeta) {
	p := a.client.Volume.NewListVolumesParams()
	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		return
	}
	for _, vol := range resp.Volumes {
		if !strings.HasPrefix(vol.Name, config.EphemeralDiskPrefix) {
			continue
		}
		a.setMetadata(config.Volume, vol.Id, &meta)
	}
}
