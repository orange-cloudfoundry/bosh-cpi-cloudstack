package action

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) HasDisk(cid apiv1.DiskCID) (bool, error) {
	volumes, err := a.findVolumesByName(cid)
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Has disk failed on disk %s", cid.AsString())
	}
	if len(volumes) > 1 {
		return false, bosherr.WrapErrorf(
			fmt.Errorf("multiple volumes found with this name"),
			"Has disk failed on disk %s", cid.AsString())
	}
	return len(volumes) == 1, nil
}
