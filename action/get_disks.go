package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) GetDisks(cid apiv1.VMCID) ([]apiv1.DiskCID, error) {
	a.logger.Info("get_disks", "fetching disks for vm '%s'...", cid.AsString())

	disks, err := a.getDisks(cid.AsString())
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not fetch disks of vm '%s'", cid.AsString())
		return []apiv1.DiskCID{}, err
	}

	cids := []apiv1.DiskCID{}
	for _, cDiskName := range disks {
		cids = append(cids, apiv1.NewDiskCID(cDiskName))
	}

	a.logger.Info("get_disks", "finished fetching disks for vm '%s'", cid.AsString())
	return cids, nil
}

func (a CPI) getDisks(vmName string) ([]string, error) {
	res := []string{}
	vm, err := a.vmFindByName(vmName)
	if err != nil {
		return nil, err
	}

	volumes, err := a.volumesFindPersistentByVM(vm)
	if err != nil {
		return nil, err
	}

	for _, cVolume := range volumes {
		res = append(res, cVolume.Name)
	}
	return res, nil
}
