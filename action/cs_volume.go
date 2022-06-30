package action

import (
	"strings"
	"time"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) volumesFindByVM(vm *cloudstack.VirtualMachine) ([]*cloudstack.Volume, error) {
	a.client.DefaultOptions()
	a.logger.Debug("volumesFindByVM", "finding volumes for vm '%s' (%s)...", vm.Name, vm.Id)

	p := a.client.Volume.NewListVolumesParams()
	p.SetVirtualmachineid(vm.Id)
	p.SetType(string(config.Datadisk))
	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not list disks of vm '%s' (%s)", vm.Id, vm.Name)
		a.logger.Error("volumesFindByVM", err.Error())
		return nil, err
	}

	a.logger.Debug("volumesFindByVM", "finished finding volumes for vm '%s' (%s)...", vm.Name, vm.Id)
	return resp.Volumes, nil
}


func (a CPI) volumesFindByTags(tags map[string]string) ([]*cloudstack.Volume, error) {
	a.client.DefaultOptions()
	a.logger.Debug("volumesFindByTags", "finding volumes for tags '%#v'...", tags)

	p := a.client.Volume.NewListVolumesParams()
	p.SetTags(tags)
	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not list disks with tags '%#v'", tags)
		a.logger.Error("volumesFindByTags", err.Error())
		return nil, err
	}

	a.logger.Debug("volumesFindByTags", "finished finding volumes for tags '%#v'...", tags)
	return resp.Volumes, nil
}

func (a CPI) volumesFindEphemeralByVM(vm *cloudstack.VirtualMachine) ([]*cloudstack.Volume, error) {
	volumes, err := a.volumesFindByVM(vm)
	if err != nil {
		return nil, err
	}
	return a.volumesFilterEphemeral(volumes), nil
}

func (a CPI) volumesFindPersistentByVM(vm *cloudstack.VirtualMachine) ([]*cloudstack.Volume, error) {
	volumes, err := a.volumesFindByVM(vm)
	if err != nil {
		return nil, err
	}
	return a.volumesFilterPersistent(volumes), nil
}

func (a CPI) volumesFilterEphemeral(volumes []*cloudstack.Volume) ([]*cloudstack.Volume) {
	res := []*cloudstack.Volume{}
	for _, cVolume := range volumes {
		if strings.HasPrefix(cVolume.Name, config.EphemeralDiskPrefix) {
			res = append(res, cVolume)
		}
	}
	return res
}

func (a CPI) volumesFilterPersistent(volumes []*cloudstack.Volume) ([]*cloudstack.Volume) {
	res := []*cloudstack.Volume{}
	for _, cVolume := range volumes {
		if strings.HasPrefix(cVolume.Name, config.PersistenceDiskPrefix) {
			res = append(res, cVolume)
		}
	}
	return res
}

func (a CPI) volumesFilterDetached(volumes []*cloudstack.Volume) ([]*cloudstack.Volume) {
	res := []*cloudstack.Volume{}
	for _, cVolume := range volumes {
		if cVolume.Vmname == "" {
			res = append(res, cVolume)
		}
	}
	return res
}

func (a CPI) volumesFilterReady(volumes []*cloudstack.Volume) ([]*cloudstack.Volume) {
	res := []*cloudstack.Volume{}
	for _, cVolume := range volumes {
		if cVolume.Destroyed == false {
			res = append(res, cVolume)
		}
	}
	return res
}

func (a CPI) volumesFilterCreatedBefore(volumes []*cloudstack.Volume, date time.Time) ([]*cloudstack.Volume) {
	res := []*cloudstack.Volume{}
	for _, cVolume := range volumes {
		created, err := time.Parse(time.RFC3339, cVolume.Created)
		if err != nil {
			err := bosherr.WrapErrorf(err, "unable to parse createdat '%s' as RFC3339, skipping", cVolume.Created)
			a.logger.Warn("volumesFilterCreatedBefore", err.Error())
			continue
		}
		if created.Before(date) {
			res = append(res, cVolume)
		}
	}
	return res
}


func (a CPI) volumesFindByName(diskCID string) ([]*cloudstack.Volume, error) {
	p := a.client.Volume.NewListVolumesParams()
	p.SetName(diskCID)
	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		return []*cloudstack.Volume{}, err
	}
	return resp.Volumes, nil
}

func (a CPI) volumeFindByName(diskCID string) (*cloudstack.Volume, error) {
	volumes, err := a.volumesFindByName(diskCID)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "could not list volumes '%s'", diskCID)
	}
	if len(volumes) > 1 {
		return nil, bosherr.Errorf("too many volumes '%s', found %d", diskCID, len(volumes))
	}
	if len(volumes) == 0 {
		return nil, bosherr.Errorf("no volume found '%s'", diskCID)
	}
	return volumes[0], nil
}

func (a CPI) volumeFindDetached(diskCID string) (*cloudstack.Volume, error) {
	volume, err := a.volumeFindByName(diskCID)
	if err != nil {
		return nil, err
	}
	if volume.Vmname != "" {
		return nil, bosherr.Errorf("volume '%s' already attached to vm '%s'", diskCID, volume.Vmname)
	}
	return volume, nil
}

func (a CPI) volumeDetach(volume *cloudstack.Volume) (error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DetachVolume)

	a.logger.Debug("volumeDetach", "detaching volume '%s' from vm '%s'...", volume.Id, volume.Virtualmachineid)
	detachParams := a.client.Volume.NewDetachVolumeParams()
	detachParams.SetId(volume.Id)
	_, err := a.client.Volume.DetachVolume(detachParams)
	if err != nil {
		err = bosherr.WrapErrorf(err, "error when detaching volume '%s' from vm '%s'", volume.Id, volume.Virtualmachineid)
		a.logger.Error("volumeDetach", "%s", err)
		return err
	}
	a.logger.Debug("volumeDetach", "finished detaching volume '%s' from vm '%s'...", volume.Id, volume.Virtualmachineid)
	return nil
}


func (a CPI) volumesDetach(volumes []*cloudstack.Volume) (error) {
	for _, cVolume := range volumes {
		if err := a.volumeDetach(cVolume); err != nil {
			return err
		}
	}
	return nil
}

func (a CPI) volumeAttach(volume *cloudstack.Volume, vm *cloudstack.VirtualMachine) (int64, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.AttachVolume)

	a.logger.Debug("volumeAttach", "attaching volume '%s' to vm '%s'...", volume.Id, vm.Id)
	p := a.client.Volume.NewAttachVolumeParams(volume.Id, vm.Id)
	resp, err := a.client.Volume.AttachVolume(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "cannot attach volume '%s' to vm '%s'", volume.Id, vm.Id)
		a.logger.Error("volumeAttach", "%s", err)
		return 0, err
	}
	a.logger.Debug("volumeAttach", "finished attaching volume '%s' to vm '%s'", volume.Id, vm.Id)
	return resp.Deviceid, nil
}

func (a CPI) volumeCreate(diskName string, size int, zone *cloudstack.Zone, offer *cloudstack.DiskOffering, cid string) (error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.CreateVolume)

	p := a.client.Volume.NewCreateVolumeParams()
	p.SetName(diskName)
	p.SetZoneid(zone.Id)
	p.SetDiskofferingid(offer.Id)
	if offer.Iscustomized {
		size = int(size / 1024)
		p.SetSize(int64(size))
	}

	a.logger.Debug("volumeCreate", "creating volume '%s' in zone '%d' and offer '%d'...", diskName, zone.Id, offer.Id)
	_, err := a.client.Volume.CreateVolume(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "cannot create volume '%s' in zone '%d' and offer '%d'", diskName, zone.Id, offer.Id)
		a.logger.Error("volumeCreate", "%s", err)
		return err
	}

	a.logger.Debug("volumeCreate", "finished creating volume '%s' in zone '%d' and offer '%d'", diskName, zone.Id, offer.Id)
	return nil
}

func (a CPI) volumeDelete(volume *cloudstack.Volume) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteVolume)

	a.logger.Debug("volumeDelete", "deleting volume '%s' (%s)...", volume.Name, volume.Id)
	p := a.client.Volume.NewDeleteVolumeParams(volume.Id)
	_, err := a.client.Volume.DeleteVolume(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not delete volume '%s' (%s)", volume.Name, volume.Id)
		a.logger.Error("volumeDelete", err.Error())
		return err
	}
	a.logger.Debug("volumeDelete", "finished deleting volume '%s' (%s)", volume.Name, volume.Id)

	return nil
}

func (a CPI) volumesDelete(volumes []*cloudstack.Volume) error {
	for _, cVolume := range volumes {
		if err := a.volumeDelete(cVolume); err != nil {
			return err
		}
	}
	return nil
}

func (a CPI) volumeResize(volume *cloudstack.Volume, newSizeMB int, switchOffer *cloudstack.DiskOffering) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.ResizeVolume)

	a.logger.Debug("volumeResize", "resizing volume '%s' (%s) to '%d MB'...", volume.Name, volume.Id, newSizeMB)
	p := a.client.Volume.NewResizeVolumeParams(volume.Id)
	p.SetSize(int64(newSizeMB / 1024))
	if switchOffer != nil {
		p.SetDiskofferingid(switchOffer.Id)
	}

	_, err := a.client.Volume.ResizeVolume(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not resize volume '%s' (%s) to ", volume.Name, volume.Id, newSizeMB)
		a.logger.Error("volumeResize", err.Error())
		return err
	}
	a.logger.Debug("volumeResize", "finished resizing volume '%s' (%s) to '%d MB'", volume.Name, volume.Id, newSizeMB)
	return nil
}

func (a CPI) volumeSnapshot(volume *cloudstack.Volume) (string, error) {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.SnapshotVolume)

	a.logger.Debug("volumeSnapshot", "snapshoting volume '%s' (%s)...", volume.Name, volume.Id)
	p := a.client.Snapshot.NewCreateSnapshotParams(volume.Id)
	resp, err := a.client.Snapshot.CreateSnapshot(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not snapshot volume '%s' (%s)", volume.Name, volume.Id)
		a.logger.Error("volumeSnapshot", err.Errorf())
		return "", err
	}
	a.logger.Debug("volumeSnapshot", "finished snapshoting volume '%s' (%s)", volume.Name, volume.Id)
	return resp.Id, nil
}
