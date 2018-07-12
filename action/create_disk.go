package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"fmt"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
)

func (a CPI) CreateDisk(size int, props apiv1.DiskCloudProps, cid *apiv1.VMCID) (apiv1.DiskCID, error) {
	var diskProps DiskCloudProperties
	err := props.As(&diskProps)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid.AsString())
	}

	diskOfferName := diskProps.DiskOffering
	if diskOfferName == "" {
		diskOfferName = a.config.CloudStack.DefaultOffers.Disk
		a.logger.Info("create_disk", "Using default disk offering %s because not set in properties", diskOfferName)
	}

	if diskOfferName == "" {
		diskOfferName = a.config.CloudStack.DefaultOffers.CustomDisk
		a.logger.Info("create_disk", "Using default custom disk offering %s because there is no default disk offers", diskOfferName)
	}

	diskName := fmt.Sprintf("%s%s", config.PersistenceDiskPrefix, uuid.NewV4().String())

	zoneId, err := a.findZoneId()
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid.AsString())
	}

	offer, err := a.findDiskOfferingByName(diskOfferName)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid.AsString())
	}

	p := a.client.Volume.NewCreateVolumeParams()
	p.SetName(diskName)
	p.SetZoneid(zoneId)
	p.SetDiskofferingid(offer.Id)

	if offer.Iscustomized {
		size = int(size / 1024)
		p.SetSize(int64(size))
	}

	_, err = a.client.Volume.CreateVolume(p)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Cannot create disk for vm %s", cid.AsString())
	}

	return apiv1.NewDiskCID(diskName), nil
}
