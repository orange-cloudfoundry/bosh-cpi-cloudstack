package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"fmt"
)

func (a CPI) HasVM(cid apiv1.VMCID) (bool, error) {
	vms, err := a.findVmsByName(cid)
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Has vm failed on vm %s", cid.AsString())
	}
	if len(vms) > 1 {
		return false, bosherr.WrapErrorf(
			fmt.Errorf("multiple vm found with this name"),
			"Has vm failed on vm %s", cid.AsString())
	}
	return len(vms) == 1, nil
}
