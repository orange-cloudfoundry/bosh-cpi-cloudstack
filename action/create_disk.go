package action

import (
	"fmt"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
)

func (a CPI) CreateDisk(size int, props apiv1.DiskCloudProps, cid *apiv1.VMCID) (apiv1.DiskCID, error) {
	cidStr := "'not used'"
	if cid != nil {
		cidStr = cid.AsString()
	}
	var diskProps DiskCloudProperties
	err := props.As(&diskProps)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cidStr)
	}

	diskOfferName := diskProps.DiskOffering
	if diskOfferName == "" {
		diskOfferName = a.config.CloudStack.DefaultOffer.Disk
		a.logger.Info("create_disk", "Using default disk offering %s because not set in properties", diskOfferName)
	}

	diskName := fmt.Sprintf("%s%s", config.PersistenceDiskPrefix, uuid.NewV4().String())

	a.logger.Info("create_disk", "Creating disk %s ...", diskName)
	_, err = a.createVolume(diskName, size, diskOfferName, cidStr)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cidStr)
	}
	a.logger.Info("create_disk", "Finished creating disk %s .", diskName)

	return apiv1.NewDiskCID(diskName), nil
}

func (a CPI) createEphemeralDisk(size int, diskProps DiskCloudProperties, cid string) (apiv1.DiskCID, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.CreateVolume)

	diskOfferName := diskProps.EphemeralDiskOffering
	if diskOfferName == "" {
		diskOfferName = a.config.CloudStack.DefaultOffer.EphemeralDisk
		a.logger.Info("create_disk", "Using default disk offering %s because not set in properties", diskOfferName)
	}

	diskName := fmt.Sprintf("%s%s", config.EphemeralDiskPrefix, uuid.NewV4().String())

	_, err := a.createVolume(diskName, size, diskOfferName, cid)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid)
	}

	// we tag ephemeral disk with director name
	// this is useful for automatic cleaning
	if a.config.CloudStack.DirectorName != "" {
		meta := apiv1.NewCloudKVs(map[string]interface{}{
			"director": a.config.CloudStack.DirectorName,
		})
		err := a.setMetadata(config.Volume, diskName, &meta)
		if err != nil {
			a.logger.Warn("create_ephemeral_disk", "Error occurred when setting metadata on ephemeral disk %s: %s", diskName, err.Error())
		}
	}

	return apiv1.NewDiskCID(diskName), nil
}

func (a CPI) createVolume(diskName string, size int, diskOfferName string, cid string) (*cloudstack.CreateVolumeResponse, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.CreateVolume)

	if diskOfferName == "" {
		diskOfferName = a.config.CloudStack.DefaultOffer.CustomDisk
		a.logger.Info("create_disk", "Using default custom disk offering %s because there is no default disk offers", diskOfferName)
	}

	zoneId, err := a.findZoneID()
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid)
	}

	offer, err := a.findDiskOfferingByName(diskOfferName)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid)
	}

	p := a.client.Volume.NewCreateVolumeParams()
	p.SetName(diskName)
	p.SetZoneid(zoneId)
	p.SetDiskofferingid(offer.Id)

	if offer.Iscustomized {
		size = int(size / 1024)
		p.SetSize(int64(size))
	}

	resp, err := a.client.Volume.CreateVolume(p)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid)
	}

	return resp, nil
}
