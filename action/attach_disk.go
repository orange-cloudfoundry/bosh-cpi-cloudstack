package action

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) AttachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	return bosherr.Errorf("attach_disk '%s' to vm '%s' is no more supported from api v1", diskCID.AsString(), vmCID.AsString())
}

func (a CPI) AttachDiskV2(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) (apiv1.DiskHint, error) {
	a.logger.Info("attach_disk", "attaching disk '%s' to vm '%s'...", diskCID.AsString(), vmCID.AsString())
	hints, err := a.attachDisk(vmCID, diskCID)
	if err != nil {
		a.logger.Info("attach_disk", bosherr.WrapErrorf(err, "error while attaching disk '%s' to vm '%s'", diskCID.AsString(), vmCID.AsString()))
		return apiv1.DiskHint{}, err
	}
	a.logger.Info("attach_disk", "finished attaching disk '%s' to vm '%s'", diskCID.AsString(), vmCID..AsString())
	return hints, err
}

func (a CPI) DetachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	a.logger.Info("detach_disk", "detaching disk '%s' from vm '%s' ...", diskCID.AsString(), vmCID.AsString())
	err = a.detachDisk(vmCID, diskCID)
	if err != nil {
		a.logger.Info("detach_disk", bosherr.WrapErrorf(err, "error while detaching disk '%s' from vm '%s'", diskCID.AsString(), vmCID.AsString())
		return err
	}
	a.logger.Info("detach_disk", "finished detaching disk '%s' from vm '%s'", diskCID.AsString(), vmCID.AsString())
}


func (a CPI) attachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) (apiv1.DiskHint, error) {
	volume, err := a.findVolumeDetached(diskCID)
	if err != nil {
		return apiv1.DiskHint{}, err
	}
	vm, err := a.findVmByName(vmCID)
	if err != nil {
		return apiv1.DiskHint{}, err
	}
	deviceID, err := a.volumeAttach(volume, vm)
	if err != nil {
		return apiv1.DiskHint{}, err
	}
	indexVol := byte('a') + byte(deviceID)
	diskHint := apiv1.NewDiskHintFromMap(map[string]interface{}{
		"path":      "/dev/xvd" + string(indexVol),
		"volumd_id": fmt.Sprintf("%d", deviceID),
	})
	return diskHint, nil
}

// DetachDisk -
// 1. if already detached do nothing
func (a CPI) detachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	if a.ctx.APIVersion == 1 {
		return bosherr.Errorf("detach_disk '%s' from vm '%s' is no more supported from api v1", diskCID, vmCID)
	}
	volume, err := a.findVolumeByName(diskCID)
	if err != nil {
		return err
	}
	// 1.
	if volume.Virtualmachineid == "" {
		return nil
	}
	return a.detachVolume(volume)
}
