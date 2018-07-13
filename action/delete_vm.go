package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

func (a CPI) DeleteVM(cid apiv1.VMCID) error {

	return nil
}

func (a CPI) deleteVMById(vmId string) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteVm)

	p := a.client.VirtualMachine.NewDestroyVirtualMachineParams(vmId)
	_, err := a.client.VirtualMachine.DestroyVirtualMachine(p)
	return err
}
