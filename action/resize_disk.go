package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"fmt"
)

func (a CPI) ResizeDisk(cid apiv1.DiskCID, newSizeMB int) error {
	a.logger.Info("resize_disk", "resizing disk '%s' to '%d MB' ...", cid.AsString(), newSizeMB)
	err := a.resizeDisk(cid.AsString(), newSizeMB)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not resize disk '%s' to new size '%d MB'", cid.AsString(), newSizeMB)
		a.logger.Error("resize_disk", err.Error())
		return err
	}
	a.logger.Info("resize_disk", "finished resizing disk '%s' to '%d MB'", cid.AsString(), newSizeMB)
	return nil
}

// resizeDisk -
// 1. we force switch to a customizable disk offer if current offer is not
func (a CPI) resizeDisk(diskName string, newSizeMB int) error {
	volume, err := a.volumeFindByName(diskName)
	if err != nil {
		return err
	}

	offer, err := a.diskOfferingByName(volume.Diskofferingname)
	if err != nil {
		return err
	}

	newSizeGB := int64(newSizeMB / 1024)
	if offer.Disksize > newSizeGB {
		msg := fmt.Errorf("disk new size < current size (current: %d GB, asked: %d GB)", offer.Disksize, newSizeGB)
		return NewNotImplementedError(msg)
	}

	// 1.
	if offer.Iscustomized {
		err = a.volumeResize(volume, newSizeMB, nil)
	} else {
		offerCustom, err := a.diskOfferingByName(a.config.CloudStack.DefaultOffer.CustomDisk)
		if err != nil {
			return err
		}
		err = a.volumeResize(volume, newSizeMB, offerCustom)
	}
	return err
}
