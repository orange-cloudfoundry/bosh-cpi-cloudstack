package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"

)

func (a CPI) CreateStemcell(imagePath string, _ apiv1.StemcellCloudProps) (apiv1.StemcellCID, error) {
	return apiv1.StemcellCID{}, nil
}
