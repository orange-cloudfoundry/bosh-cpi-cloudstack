package action

import (
	"encoding/json"
	"fmt"

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
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Printf("error while unmarshalling meta: %v", err)
	}
	return data
}
