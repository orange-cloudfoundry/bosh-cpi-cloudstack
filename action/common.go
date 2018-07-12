package action

import (
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

func (a CPI) setMetadata(tagType config.Tags, cid string, meta util.MetaMarshal) error {
	params := a.client.Resourcetags.NewCreateTagsParams([]string{cid}, string(tagType), util.ConvertMapToTags(meta))
	_, err := a.client.Resourcetags.CreateTags(params)
	if err != nil {
		return bosherr.WrapErrorf(err, "Setting %s metadata '%s'", tagType, cid)
	}
	return nil
}

func (a CPI) findVmId(cid apiv1.VMCID) (string, error) {
	p := a.client.VirtualMachine.NewListVirtualMachinesParams()
	p.SetName(cid.AsString())
	resp, err := a.client.VirtualMachine.ListVirtualMachines(p)
	if err != nil {
		return "", bosherr.WrapErrorf(err, "Can't find vm name '%s'", cid.AsString())
	}
	if len(resp.VirtualMachines) == 0 {
		return "", bosherr.Errorf("Can't find vm name '%s'", cid.AsString())

	}
	return resp.VirtualMachines[0].Id, nil
}

func (a CPI) findVolumeId(cid apiv1.DiskCID) (string, error) {
	p := a.client.Volume.NewListVolumesParams()
	p.SetName(cid.AsString())
	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		return "", bosherr.WrapErrorf(err, "Can't find disk name '%s'", cid.AsString())
	}
	if len(resp.Volumes) == 0 {
		return "", bosherr.Errorf("Can't find disk name '%s'", cid.AsString())

	}
	return resp.Volumes[0].Id, nil
}
