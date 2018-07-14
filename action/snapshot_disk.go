package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"strings"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

/*
logger.info("snapshot disk");
//TODO: only for persistent disk
String csDiskId = api.getVolumeApi().getVolume(disk_id).getId();
AsyncCreateResponse async = api.getSnapshotApi().createSnapshot(csDiskId, CreateSnapshotOptions.Builder.domainId("domain"));

jobComplete = retry(new JobComplete(api), 1200, 3, 5, SECONDS);
jobComplete.apply(async.getJobId());

//FIXME
return null;
 */
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

	p := a.client.Snapshot.NewCreateSnapshotParams(volume.Id)
	resp, err := a.client.Snapshot.CreateSnapshot(p)
	if err != nil {
		return apiv1.SnapshotCID{}, bosherr.WrapErrorf(err, "Could not create snapshot for disk %s", diskCID.AsString())
	}
	a.setMetadata(config.Snapshot, resp.Id, &meta)

	return apiv1.NewSnapshotCID(resp.Id), nil
}
