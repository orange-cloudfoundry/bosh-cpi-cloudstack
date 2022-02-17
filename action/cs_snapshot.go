package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)


func (a CPI) snapshotDelete(ID string) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteSnapshotVolume)

	a.logger.Debug("snapshotDelete", "deleting snapshot '%s' ...", ID)
	p := a.client.Snapshot.NewDeleteSnapshotParams(ID)
	_, err := a.client.Snapshot.DeleteSnapshot(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not delete snapshot '%s'", ID)
		a.logger.Error("snapshotDelete", err.Error())
		return err
	}
	a.logger.Debug("delete_disk", "Finished deleting snapshot %s .", ID)
	return nil
}
