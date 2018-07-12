package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"strings"
)

func (a CPI) GetDisks(cid apiv1.VMCID) ([]apiv1.DiskCID, error) {
	diskCids := make([]apiv1.DiskCID, 0)
	vm, err := a.findVmByName(cid)
	if err != nil {
		return diskCids, bosherr.WrapErrorf(err, "Cannot getting disks for vm %s", cid.AsString())
	}

	p := a.client.Volume.NewListVolumesParams()
	p.SetVirtualmachineid(vm.Id)
	p.SetType(string(config.Datadisk))
	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		return diskCids, bosherr.WrapErrorf(err, "Cannot getting disks for vm %s", cid.AsString())
	}

	for _, vol := range resp.Volumes {
		if !strings.HasPrefix(vol.Name, config.PersistenceDiskPrefix) {
			continue
		}
		diskCids = append(diskCids, apiv1.NewDiskCID(vol.Name))
	}
	return diskCids, nil
}
