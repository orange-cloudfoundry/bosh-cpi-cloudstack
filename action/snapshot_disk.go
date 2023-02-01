package action

import (
	"strings"

	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SnapshotDisk(diskCID apiv1.DiskCID, meta apiv1.DiskMeta) (apiv1.SnapshotCID, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.SnapshotVolume)

	volumes, err := a.findVolumesByName(diskCID)
	if err != nil {
		return apiv1.SnapshotCID{}, bosherr.WrapErrorf(err, "Error when finding disk %s", diskCID.AsString())
	}

	if len(volumes) > 1 {
		return apiv1.SnapshotCID{}, bosherr.Errorf("Too much volume with name %s", diskCID.AsString())
	}

	if len(volumes) == 0 {
		return apiv1.SnapshotCID{}, bosherr.Errorf("No volume found with name %s", diskCID.AsString())
	}
	volume := volumes[0]
	if strings.HasPrefix(volume.Name, config.PersistenceDiskPrefix) {
		return apiv1.SnapshotCID{}, bosherr.Errorf("Volume found with name %s is not a persistent disk", diskCID.AsString())
	}

	a.logger.Info("resize_disk", "Snapshotting disk %s ...", diskCID.AsString())
	p := a.client.Snapshot.NewCreateSnapshotParams(volume.Id)
	resp, err := a.client.Snapshot.CreateSnapshot(p)
	if err != nil {
		return apiv1.SnapshotCID{}, bosherr.WrapErrorf(err, "Could not create snapshot for disk %s", diskCID.AsString())
	}
	a.logger.Info("resize_disk", "Finished snapshotting disk %s .", diskCID.AsString())

	err = a.setMetadata(config.Snapshot, resp.Id, &meta)
	if err != nil {
		a.logger.Warn("set_metadata", "an error occurred while setting metadata to: %s = %v", resp.Id, &meta)
	}

	return apiv1.NewSnapshotCID(resp.Id), nil
}
