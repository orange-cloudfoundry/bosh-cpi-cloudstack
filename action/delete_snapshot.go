package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) DeleteSnapshot(cid apiv1.SnapshotCID) error {
	a.logger.Info("delete_snapshot", "deleting snapshot '%s' ...", cid.AsString())

	err := a.snapshotDelete(cid.AsString())
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not delete snapshot '%s'", cid.AsString())
		a.logger.Error("delete_snapshot", err.Error())
		return err
	}

	a.logger.Info("delete_snapshot", "finished deleting snapshot '%s'", cid.AsString())
	return nil
}
