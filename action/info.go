package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

func (a CPI) Info() (apiv1.Info, error) {
	return apiv1.Info{StemcellFormats: []string{"cloudstack-vhdx"}}, nil
}
