package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

func (a CPI) RebootVM(_ apiv1.VMCID) error {
	return nil
}
