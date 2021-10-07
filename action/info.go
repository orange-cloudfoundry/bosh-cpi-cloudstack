package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
)

func (a CPI) Info() (apiv1.Info, error) {
	return apiv1.Info{
		APIVersion:      2,
		StemcellFormats: []string{"cloudstack-vhdx"},
	}, nil
}
