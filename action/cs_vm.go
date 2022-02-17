package action

import (
	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) vmsFindByName(name string) ([]*cloudstack.VirtualMachine, error) {
	a.client.DefaultOptions()
	a.logger.Debug("vmsFindByName", "listing virtual machine with name '%s'...", name)

	p := a.client.VirtualMachine.NewListVirtualMachinesParams()
	p.SetName(name)
	resp, err := a.client.VirtualMachine.ListVirtualMachines(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "too many virtual machines with name '%s'", name)
		a.logger.Errof("vmFindByName", err.Error())
		return nil, err
	}

	a.logger.Debug("vmsFindByName", "finished listing virtual machine with name '%s'", name)
	return resp.VirtualMachines, nil
}

func (a CPI) vmFindByName(name string) (*cloudstack.VirtualMachine, error) {
	resp, err := a.vmsFindByName(name)
	if err != nil {
		return nil, err
	}

	if len(resp) > 1 {
		err = bosherr.Errorf("too many virtual machines with name '%s'", name)
		a.logger.Errof("vmFindByName", err.Error())
		return nil, err
	}

	if len(resp) == 0 {
		err = bosherr.Errorf("could not find many virtual machine with name '%s'", name)
		a.logger.Errof("vmFindByName", err.Error())
		return nil, err
	}

	return resp[0], nil
}

func (a CPI) vmStop(vm *cloudstack.VirtualMachine) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.StopVm)
	a.logger.Debug("vmStop", "stopping vm '%s' (%s)...", vm.Name, vm.Id)

	p := a.client.VirtualMachine.NewStopVirtualMachineParams(vm.Id)
	p.SetForced(true)
	_, err := a.client.VirtualMachine.StopVirtualMachine(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "error when stopping vm '%s' (%s)'", vm.Id, vm.Name)
		a.logger.Error("vmStop", "%s", err)
		return err
	}

	a.logger.Debug("vmStop", "finished stopping vm '%s' (%s)", vm.Name, vm.Id)
	return nil
}


func (a CPI) vmDestroy(vm *cloudstack.VirtualMachine) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteVm)
	a.logger.Debug("vmDestroy", "destroying vm '%s' (%s)...", vm.Name, vm.Id)

	p := a.client.VirtualMachine.NewDestroyVirtualMachineParams(vm.Id)
	p.SetExpunge(a.config.CloudStack.ExpungeVm)
	_, err := a.client.VirtualMachine.DestroyVirtualMachine(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "error when destroying vm '%s' (%s)'", vm.Id, vm.Name)
		a.logger.Error("vmDestroy", "%s", err)
		return err
	}

	a.logger.Debug("vmDestroy", "finished destroying vm '%s' (%s)", vm.Name, vm.Id)
	return nil
}

func (a CPI) vmReboot(vm *cloudstack.VirtualMachine) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.RebootVm)
	a.logger.Debug("vmReboot", "rebooting vm '%s' (%s)...", vm.Name, vm.Id)

	p := a.client.VirtualMachine.NewRebootVirtualMachineParams(vm.Id)
	_, err := a.client.VirtualMachine.RebootVirtualMachine(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "error when rebooting vm '%s' (%s)'", vm.Id, vm.Name)
		a.logger.Error("vmReboot", "%s", err)
		return err
	}

	a.logger.Debug("vmReboot", "finished rebooting vm '%s' (%s)", vm.Name, vm.Id)
	return nil
}

