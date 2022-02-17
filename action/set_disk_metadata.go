package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) SetDiskMetadata(cid apiv1.DiskCID, meta apiv1.DiskMeta) error {
	a.logger.Info("set_disk_metadata", "setting metadata on disk '%s'...", cid.AsString())
	err := a.setDiskMetadata(cid.AsString(), &meta)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not set metadata on disk '%s'", cid.AsString())
		a.logger.Error("set_disk_metadata", err.Error())
		return err
	}
	a.logger.Info("set_disk_metadata", "finished setting metadata on disk '%s'...", cid.AsString())
	return nil
}


// setDiskMetadata -
// 1. delete any tag that are already present
func (a CPI) setDiskMetadata(diskName, meta util.MetaMarshal) error {
	volume, err := a.volumeFindByName(diskName)
	if err != nil {
		return err
	}

	// 1.
	tags, err := a.tagList(string(config.Volume), volume.ID)
	if (err == nil) && (len(tags) != 0) {
		err = a.tagDelete(string(config.Volume), volume.ID)
		if err != nil {
			return err
		}
	}

	tagMap := util.ConvertMapToTags(meta)
	tagMap["director_uuid"] = a.ctx.DirectorUUID
	return a.tagCreate(string(config.Volume), volume.ID, tagMap)
}
