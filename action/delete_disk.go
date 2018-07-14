package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"fmt"
)

func (a CPI) DeleteDisk(cid apiv1.DiskCID) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteVolume)

	volumes, err := a.findVolumesByName(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Cannot delete disk %s", cid.AsString())
	}
	if len(volumes) > 1 {
		return bosherr.WrapErrorf(
			fmt.Errorf("multiple volumes found with this name"),
			"Cannot delete disk %s", cid.AsString())
	}
	if len(volumes) == 0 {
		return nil
	}

	a.logger.Info("delete_disk", "Deleting disk %s ...", cid.AsString())
	p := a.client.Volume.NewDeleteVolumeParams(volumes[0].Id)
	_, err = a.client.Volume.DeleteVolume(p)
	if err != nil {
		return bosherr.WrapErrorf(err, "Cannot delete disk %s", cid.AsString())
	}
	a.logger.Info("delete_disk", "Finished deleting disk %s .", cid.AsString())
	return nil
}
