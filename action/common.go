package action

import (
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/util"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) setMetadata(tagType config.Tags, cid string, meta util.MetaMarshal) error {
	params := a.client.Resourcetags.NewCreateTagsParams([]string{cid}, string(tagType), util.ConvertMapToTags(meta))
	_, err := a.client.Resourcetags.CreateTags(params)
	if err != nil {
		return bosherr.WrapErrorf(err, "Setting %s metadata '%s'", tagType, cid)
	}
	return nil
}
