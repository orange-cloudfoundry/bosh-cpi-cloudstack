package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
)

func (a CPI) DeleteStemcell(cid apiv1.StemcellCID) error {
	a.logger.Info("delete_stemcell", "deleting stemcell %s ...", cid.AsString())
	err := a.deleteStemcell(cid.AsString())
	if err != nil {
		err = bosherr.WrapErrorf(err, "unable to delete stemcell '%s'", cid.AsString())
		a.logger.Error("delete_stemcell", err.Error())
	}
	a.logger.Info("delete_stemcell", "finished deleting stemcell %s", cid.AsString())
	return nil
}

func (a CPI) deleteStemcell(stemcell string) error {
	zone, err := a.zoneFindDefault()
	if err != nil {
		return err
	}
	template, err := a.templateFindByName(stemcell, zone)
	if err != nil {
		return err
	}
	err = a.templateDelete(template, zone)
	if err != nil {
		return err
	}
	return nil
}
