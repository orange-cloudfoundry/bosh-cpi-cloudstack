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


func (a CPI) setDiskMetadata(diskName string, meta util.MetaMarshal) error {
	volume, err := a.volumeFindByName(diskName)
	if err != nil {
		return err
	}
	return a.setMetadata(string(config.Volume), volume.Id, meta)
}
