package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) HasDisk(cid apiv1.DiskCID) (bool, error) {
	a.logger.Info("has_disk", "checking disk exists for '%s'...", cid.AsString())

	found, err := a.hasDisk(cid.AsString())
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not check if disk exists for '%s'", cid.AsString())
		a.logger.Error("has_disk", err.Error())
		return false, err
	}

	a.logger.Info("has_disk", "finished checking disk exists for '%s' (%t)", cid.AsString(), found)
	return found, nil
}

func (a CPI) hasDisk(diskName string) (bool, error) {
	volumes, err := a.volumeFindByName(diskName)
	if err != nil {
		return false, err
	}
	if len(volumes) > 1 {
		return false, bosherr.Errorf("found multiple instances of disk '%s'", diskName)
	}
	return len(volumes) == 1, nil
}
