package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) DeleteVM(cid apiv1.VMCID) error {
	a.logger.Info("delete_vm", "deleting vm '%s'...", cid.AsString())

	err := a.deleteVM(cid.AsString())
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not delete vm '%s'", cid.AsString())
		a.logger.Error("delete_stemcell", err.Error())
		return err
	}

	a.logger.Info("delete_vm", "finished deleting vm '%s'", cid.AsString())
	return nil
}


func (a CPI) deleteVM(vmID string) error {
	vm, err := a.vmFindByName(vmID)
	if err != nil {
		return err
	}

	ifrs, err := a.natIFRFindByVM(vm)
	if err != nil {
		return err
	}
	err = a.natIFRsDelete(ifrs)
	if err != nil {
		return err
	}

	ips, err := a.addressListPublicIPsByVM(vm)
	if err != nil {
		return err
	}
	err = a.natDisableForIps(ips)
	if err != nil {
		return err
	}

	err = a.vmStop(vm)
	if err != nil {
		return err
	}

	ephemeralVolumes, err := a.volumesFindEphemeralByVM(vm)
	if err != nil {
		return err
	}

	volumes, err := a.volumesFindEphemeralByVM(vm)
	if err != nil {
		return err
	}

	err = a.volumesDetach(volumes)
	if err != nil {
		return err
	}

	err = a.vmDestroy(vm)
	if err != nil {
		return err
	}

	err = a.volumesDelete(ephemeralVolumes)
	if err != nil {
		return err
	}

	return nil
}
