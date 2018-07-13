package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"fmt"
)

func (a CPI) ResizeDisk(cid apiv1.DiskCID, size int) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.ResizeVolume)

	volume, err := a.findVolumeByName(cid)
	if err != nil {
		return NewNotImplementedError(bosherr.WrapErrorf(err, "Cannot resize disk %s", cid.AsString()))
	}

	offer, err := a.findDiskOfferingByName(volume.Diskofferingname)
	if err != nil {
		return NewNotImplementedError(bosherr.WrapErrorf(err, "Cannot resize disk %s", cid.AsString()))
	}

	if offer.Disksize > int64(size/1024) {
		return NewNotImplementedError(
			fmt.Errorf(
				"Disk size requested is smaller than current disk size (current: %sGb, asked: %dGb)",
				offer.Disksize,
				int64(size/1024),
			))
	}

	p := a.client.Volume.NewResizeVolumeParams(volume.Id)
	p.SetSize(int64(size / 1024))
	if !offer.Iscustomized {
		offerCustom, err := a.findDiskOfferingByName(a.config.CloudStack.DefaultOffer.CustomDisk)
		if err != nil {
			return NewNotImplementedError(bosherr.WrapErrorf(err, "Cannot resize disk %s", cid.AsString()))
		}
		p.SetDiskofferingid(offerCustom.Id)
	}

	_, err = a.client.Volume.ResizeVolume(p)
	if err != nil {
		return NewNotImplementedError(bosherr.WrapErrorf(err, "Cannot resize disk %s", cid.AsString()))
	}
	return nil
}
