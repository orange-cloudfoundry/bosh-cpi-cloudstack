package action

import "github.com/cppforlife/bosh-cpi-go/apiv1"

func (a CPI) SnapshotDisk(apiv1.DiskCID, apiv1.DiskMeta) (apiv1.SnapshotCID, error) {
	return apiv1.SnapshotCID{}, nil
}
