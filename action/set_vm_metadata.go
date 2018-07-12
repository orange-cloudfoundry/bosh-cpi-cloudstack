package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) SetVMMetadata(cid apiv1.VMCID, meta apiv1.VMMeta) error {
	params := a.client.Resourcetags.NewCreateTagsParams([]string{cid.AsString()}, string(config.UserVm), util.ConvertMapToTags(meta))
	_, err := a.client.Resourcetags.CreateTags(params)
	if err != nil {
		return bosherr.WrapErrorf(err, "Setting vm metadata '%s'", cid.AsString())
	}
	return nil
}
