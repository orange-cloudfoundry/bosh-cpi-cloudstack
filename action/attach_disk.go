package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"strings"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/xanzy/go-cloudstack/cloudstack"
	"fmt"
)

func (a CPI) AttachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.AttachVolume)

	volumes, err := a.findVolumesByName(diskCID)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when finding disk %s on vm %s", diskCID.AsString(), vmCID.AsString())
	}

	if len(volumes) > 1 {
		return bosherr.Errorf("Too much volume with name %s", diskCID.AsString())
	}

	if len(volumes) == 0 {
		return bosherr.Errorf("No volume found with name %s", diskCID.AsString())
	}

	volume := volumes[0]
	if volume.Vmname != "" {
		return bosherr.Errorf("Volume with name %s already attached to vm %s", diskCID.AsString(), volume.Vmname)
	}

	vm, err := a.findVmByName(vmCID)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when finding vm %s", vmCID.AsString())
	}

	a.logger.Info("attach_disk", "Attaching disk %s ...", diskCID.AsString())
	p := a.client.Volume.NewAttachVolumeParams(volume.Id, vm.Id)
	resp, err := a.client.Volume.AttachVolume(p)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when attaching volume %s to vm %s", diskCID.AsString(), vmCID.AsString())
	}
	a.logger.Info("attach_disk", "Finished attaching disk %s .", diskCID.AsString())

	// we skip registry registering if disk is an ephemeral one
	if strings.HasPrefix(volume.Name, config.EphemeralDiskPrefix) {
		return nil
	}

	a.logger.Info("attach_disk", "Registering disk %s to registry ...", diskCID.AsString())
	err = a.registerDisk(vmCID, diskCID, resp)
	if err == nil {
		a.logger.Info("attach_disk", "Finished registering disk %s to registry.", diskCID.AsString())
		return nil
	}

	detachParams := a.client.Volume.NewDetachVolumeParams()
	detachParams.SetId(volume.Id)
	a.client.Volume.DetachVolume(detachParams)
	return err
}

func (a CPI) registerDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID, volumeResp *cloudstack.AttachVolumeResponse) error {
	nvSvc := a.regFactory.Create(vmCID)
	agentEnv, err := nvSvc.Fetch()
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when fetching registry for vm %s", vmCID.AsString())
	}
	indexVol := byte('a') + byte(volumeResp.Deviceid)
	agentEnv.AttachPersistentDisk(diskCID, struct {
		Path     string `json:"path"`
		VolumeId string `json:"volume_id"`
	}{
		Path:     "/dev/xvd" + string(indexVol),
		VolumeId: fmt.Sprintf("%d", volumeResp.Deviceid),
	})

	err = nvSvc.Update(agentEnv)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when updating registry for vm %s", vmCID.AsString())
	}
	return nil
}
