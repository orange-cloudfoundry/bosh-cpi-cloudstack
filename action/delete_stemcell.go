package action

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

// DeleteStemcell - Delete CS template matching given stemcell name
func (a CPI) DeleteStemcell(cid apiv1.StemcellCID) error {
	zoneid, err := a.findZoneID()
	if err != nil {
		return err
	}

	template, _, err := a.client.Template.GetTemplateByName(cid.AsString(), "executable", zoneid)
	if err != nil {
		return nil
	}

	deleteP := a.client.Template.NewDeleteTemplateParams(template.Id)
	deleteP.SetZoneid(zoneid)
	//deleteP.SetForced(true)

	a.logger.Info("delete_stemcell", "deleting template %s (%s)", template.Id, cid.AsString())
	_, err = a.client.Template.DeleteTemplate(deleteP)
	if err != nil {
		return bosherr.WrapErrorf(err, "[delete_stemcell] could not delete template %s (%s)", template.Id, cid.AsString())
	}

	a.logger.Info("delete_stemcell", "delete_stemcell success : template %s (%s)", template.Id, cid.AsString())
	return nil
}
