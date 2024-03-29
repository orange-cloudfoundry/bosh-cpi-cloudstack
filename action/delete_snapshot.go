package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) DeleteSnapshot(cid apiv1.SnapshotCID) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteSnapshotVolume)

	a.logger.Info("delete_disk", "Deleting snapshot %s ...", cid.AsString())
	p := a.client.Snapshot.NewDeleteSnapshotParams(cid.AsString())
	_, err := a.client.Snapshot.DeleteSnapshot(p)
	if err != nil {
		return bosherr.WrapErrorf(err, "Could not delete snapshot %s", cid.AsString())
	}
	a.logger.Info("delete_disk", "Finished deleting snapshot %s .", cid.AsString())
	return nil
}
