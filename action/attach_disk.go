package action

import (
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/orange-cloudfoundry/go-cloudstack/cloudstack"
	"strings"
)

func (a CPI) AttachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	_, err := a.AttachDiskV2(vmCID, diskCID)
	return err
}

func (a CPI) AttachDiskV2(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) (apiv1.DiskHint, error) {
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

	// we skip registry registering if disk is an ephemeral one
	if strings.HasPrefix(volume.Name, config.EphemeralDiskPrefix) {
		return apiv1.DiskHint{}, nil
	}

	a.logger.Info("attach_disk", "Registering disk %s to registry ...", diskCID.AsString())
	hint, err := a.registerDisk(vmCID, diskCID, resp)
	if err != nil {
		detachParams := a.client.Volume.NewDetachVolumeParams()
		detachParams.SetId(volume.Id)
		a.client.Volume.DetachVolume(detachParams)
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "unable to register disk into registry")
	}

	a.logger.Info("attach_disk", "Finished registering disk %s to registry.", diskCID.AsString())
	return hint, nil
}

func (a CPI) registerDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID, volumeResp *cloudstack.AttachVolumeResponse) (apiv1.DiskHint, error) {
	nvSvc := a.regFactory.Create(vmCID)
	agentEnv, err := nvSvc.Fetch()
	if err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Error when fetching registry for vm %s", vmCID.AsString())
	}
	indexVol := byte('a') + byte(volumeResp.Deviceid)
	diskHint := apiv1.NewDiskHintFromMap(map[string]interface{}{
		"path":      "/dev/xvd" + string(indexVol),
		"volumd_id": fmt.Sprintf("%d", volumeResp.Deviceid),
	})
	agentEnv.AttachPersistentDisk(diskCID, diskHint)
	if err = nvSvc.Update(agentEnv); err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Error when updating registry for vm %s", vmCID.AsString())
	}
	return diskHint, nil
}
