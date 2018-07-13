package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) DeleteSnapshot(cid apiv1.SnapshotCID) error {
	p := a.client.Snapshot.NewDeleteSnapshotParams(cid.AsString())
	_, err := a.client.Snapshot.DeleteSnapshot(p)
	if err != nil {
		return bosherr.WrapErrorf(err, "Could not delete snapshot %s", cid.AsString())
	}
	return nil
}
