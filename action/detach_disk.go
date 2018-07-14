package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"strings"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) DetachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DetachVolume)

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
	// if already detached do nothing
	if volume.Virtualmachineid == "" {
		return nil
	}

	a.logger.Info("detach_disk", "Detaching disk %s ...", diskCID.AsString())
	detachParams := a.client.Volume.NewDetachVolumeParams()
	detachParams.SetId(volume.Id)
	_, err = a.client.Volume.DetachVolume(detachParams)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when detaching volume %s to vm %s", diskCID.AsString(), vmCID.AsString())
	}
	a.logger.Info("detach_disk", "Finished detaching disk %s .", diskCID.AsString())

	// we skip registry registering if disk is an ephemeral one
	if strings.HasPrefix(volume.Name, config.EphemeralDiskPrefix) {
		return nil
	}

	a.logger.Info("detach_disk", "Removing disk %s from registry ...", diskCID.AsString())
	err = a.unregisterDisk(vmCID, diskCID)
	if err == nil {
		a.logger.Info("detach_disk", "Finished removing disk %s from registry.", diskCID.AsString())
		return nil
	}

	p := a.client.Volume.NewAttachVolumeParams(volume.Id, volume.Virtualmachineid)
	a.client.Volume.AttachVolume(p)
	return err
}

func (a CPI) unregisterDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	nvSvc := a.regFactory.Create(vmCID)
	agentEnv, err := nvSvc.Fetch()
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when fetching registry for vm %s", vmCID.AsString())
	}
	agentEnv.DetachPersistentDisk(diskCID)
	err = nvSvc.Update(agentEnv)
	if err != nil {
		return bosherr.WrapErrorf(err, "Error when updating registry for vm %s", vmCID.AsString())
	}
	return nil
}
