package action

import (
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
)

func (a CPI) CreateDisk(size int, props apiv1.DiskCloudProps, cid *apiv1.VMCID) (apiv1.DiskCID, error) {
	cidStr := "not-used"
	if cid != nil {
		cidStr = cid.AsString()
	}

	a.logger.Info("create_disk", "creating disk for vm '%s' ...", cidStr)
	diskName, err := a.createPersistentDisk(size, props, cidStr)
	if err != nil {
		return apiv1.DiskCID{}, err
	}
	a.logger.Info("create_disk", "finished creating disk for vm '%s', disk is '%s'", diskName)
	return apiv1.NewDiskCID(diskName), nil
}


func (a CPI) createPersistentDisk(size int, props apiv1.DiskCloudProps, cid string) (string, error) {
	var diskProps DiskCloudProperties

	diskName := fmt.Sprintf("%s%s", config.PersistenceDiskPrefix, uuid.NewV4().String())
	err := props.As(&diskProps)
	if err != nil {
		return diskName, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid)
	}
	offerName := diskProps.GetDiskOffering(a.config.CloudStack.DefaultOffer)

	err = a.createDisk(diskName, size, offerName, cid)
	return diskName, err
}


// createEphemeralDisk -
// 1. we tag ephemeral disk with director name, this is useful for automatic cleaning
func (a CPI) createEphemeralDisk(size int, diskProps DiskCloudProperties, cid string) (apiv1.DiskCID, error) {
	diskName := fmt.Sprintf("%s%s", config.EphemeralDiskPrefix, uuid.NewV4().String())
	offerName := diskProps.GetEphemeralDiskOffering(a.config.CloudStack.DefaultOffer)

	err := a.createDisk(diskName, size, offerName, cid)
	if err != nil {
		return apiv1.DiskCID{}, err
	}

	// 1.
	err = a.tagDisk(diskName)
	if err != nil {
		err = bosherr.WrapErrorf(err, "warning: could not set metadata on ephemeral disk '%s'", diskName)
		a.logger.Warn("create_disk", "%s", err)
	}
	return apiv1.NewDiskCID(diskName), nil
}

func (a CPI) createDisk(diskName string, size int, offerName string, cid string) (error) {
	zone, err := a.zoneFindDefault()
	if err != nil {
		return err
	}
	offer, err := a.diskOfferingFindByName(offerName)
	if err != nil {
		return err
	}
	err = a.volumeCreate(diskName, size, zone, offer, cid)
	if err != nil {
		return err
	}
	return nil
}

func (a CPI) tagDisk(diskName string) (error) {
	if a.config.CloudStack.DirectorName == "" {
		return nil
	}
	meta := apiv1.NewCloudKVs(map[string]interface{}{
		"director": a.config.CloudStack.DirectorName,
	})
	return a.setMetadata(config.Volume, diskName, &meta)
}
