package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"

)

func (a CPI) HasVM(cid apiv1.VMCID) (bool, error) {
	return false, nil
}
