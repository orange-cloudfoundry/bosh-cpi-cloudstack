package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"

)

func (a CPI) CreateDisk(size int, _ apiv1.DiskCloudProps, _ *apiv1.VMCID) (apiv1.DiskCID, error) {

	return apiv1.DiskCID{}, nil
}
