package action

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) AttachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	_, err := a.AttachDiskBase(vmCID, diskCID, false)
	return err
}

func (a CPI) AttachDiskV2(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) (apiv1.DiskHint, error) {
	diskHints, err := a.AttachDiskBase(vmCID, diskCID, true)
	return diskHints, err
}

func (a CPI) AttachDiskBase(vmCID apiv1.VMCID, diskCID apiv1.DiskCID, isV2 bool) (apiv1.DiskHint, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.AttachVolume)

	volumes, err := a.findVolumesByName(diskCID)
	if err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Error when finding disk %s on vm %s", diskCID.AsString(), vmCID.AsString())
	}

	if len(volumes) > 1 {
		return apiv1.DiskHint{}, bosherr.Errorf("Too much volume with name %s", diskCID.AsString())
	}

	if len(volumes) == 0 {
		return apiv1.DiskHint{}, bosherr.Errorf("No volume found with name %s", diskCID.AsString())
	}

	volume := volumes[0]
	if volume.Vmname != "" {
		return apiv1.DiskHint{}, bosherr.Errorf("Volume with name %s already attached to vm %s", diskCID.AsString(), volume.Vmname)
	}

	vm, err := a.findVmByName(vmCID)
	if err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Error when finding vm %s", vmCID.AsString())
	}

	a.logger.Info("attach_disk", "Attaching disk %s ...", diskCID.AsString())
	p := a.client.Volume.NewAttachVolumeParams(volume.Id, vm.Id)
	resp, err := a.client.Volume.AttachVolume(p)
	if err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Error when attaching volume %s to vm %s", diskCID.AsString(), vmCID.AsString())
	}
	a.logger.Info("attach_disk", "Finished attaching disk %s .", diskCID.AsString())

	indexVol := byte('a') + byte(resp.Deviceid)
	diskHint := apiv1.NewDiskHintFromMap(map[string]interface{}{
		"path":      "/dev/xvd" + string(indexVol),
		"volume_id": fmt.Sprintf("%d", resp.Deviceid),
	})

	if !isV2 {
		a.logger.Info("attach_disk", "Registering disk %s to registry ...", diskCID.AsString())
		if err := a.registerDisk(volume.Name, vmCID, diskCID, diskHint); err != nil {
			detachParams := a.client.Volume.NewDetachVolumeParams()
			detachParams.SetId(volume.Id)
			if _, err := a.client.Volume.DetachVolume(detachParams); err != nil {
				a.logger.Error("attach_disk", "error while detaching volume with detachParams: %s: %s", detachParams, err)
			}
			return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "unable to register disk into registry")
		}
		a.logger.Info("attach_disk", "Finished registering disk %s to registry.", diskCID.AsString())
	}

	return diskHint, nil
}

func (a CPI) registerDisk(name string, vmCID apiv1.VMCID, diskCID apiv1.DiskCID, hint apiv1.DiskHint) error {
	// we skip registry registering if disk is an ephemeral one
	if strings.HasPrefix(name, config.EphemeralDiskPrefix) {
		a.logger.Debug("attach_disk", "skip registering ephemeral disks ...")
		return nil
	}
	nvSvc := a.regFactory.Create(vmCID)
	agentEnv, err := nvSvc.Fetch()
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when fetching registry for vm %s", vmCID.AsString())
	}
	agentEnv.AttachPersistentDisk(diskCID, hint)
	if err = nvSvc.Update(agentEnv); err != nil {
		return bosherr.WrapErrorf(err, "Error when updating registry for vm %s", vmCID.AsString())
	}
	return nil
}
