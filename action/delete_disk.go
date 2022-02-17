package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) DeleteDisk(cid apiv1.DiskCID) error {
	a.logger.Info("delete_disk", "deleting disk '%s' ...", cid.AsString())

	err := a.deleteDisk(cid.AsString())
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not delete disk '%s'", cid.AsString())
		a.logger.Error("delete_disk", err.Error())
		return err
	}
	a.logger.Info("delete_disk", "finished deleting disk '%s'", cid.AsString())
	return nil
}

func (a CPI) deleteDisk(cid string) error {
	volumes, err := a.volumesFindByName(cid)
	if err != nil {
		return err
	}
	if len(volumes) > 1 {
		return bosherr.Errorf("too many volumes '%s', found %d", cid, len(volumes))
	}
	if len(volumes) == 0 {
		return nil
	}
	return a.volumeDelete(volumes[0])
}
