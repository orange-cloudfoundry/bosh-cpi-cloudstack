package action

import (
	"encoding/json"

	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
)

type VMEnv struct {
	Bosh BoshEnv `json:"bosh"`
}

type BoshEnv struct {
	Group  string   `json:"group"`
	Groups []string `json:"groups"`
}

func NewVMEnv(vmEnv apiv1.VMEnv) VMEnv {
	var data VMEnv
	b, _ := vmEnv.MarshalJSON()
	json.Unmarshal(b, &data)
	return data
}
