package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"strings"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SnapshotDisk(diskCID apiv1.DiskCID, meta apiv1.DiskMeta) (apiv1.SnapshotCID, error) {
	a.logger.Info("snapshot_disk", "snapshoting disk '%s' ...", diskCID.AsString())
	id, err := a.snapshotDisk(diskCID.AsString(), &meta)
	if err != nil {
		err = bosherr.WrapError(err, "could not snapshot disk '%s'", diskCID.AsString()))
		a.logger.Error("snapshot_disk", err.Error())
	}
	a.logger.Info("snapshot_disk", "finished snapshoting disk '%s'", diskCID.AsString())
	return apiv1.NewSnapshotCID(id), nil
}

func (a CPI) snapshotDisk(diskName string, meta util.MetaMarshal) (string, error) {
	volume, err := a.volumeFindByName(diskName)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(volume.Name, config.PersistenceDiskPrefix) {
		return "", bosherr.Errorf("could not snapshot non persistent volume '%s' (%s)", volume.Name, volume.Id)
	}

	id, err := a.volumeSnapshot(volume)
	if err != nil {
		return "", err
	}

	err = a.setMetadata(string(config.Snapshot), id, meta)
	if err != nil {
		return "", err
	}

	return id, nil
}
