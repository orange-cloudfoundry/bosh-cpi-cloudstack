package action

import (
	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)


func (a CPI) osTypesFindByDescr(descr string) ([]*cloudstack.OsType, error) {
	a.logger.Debug("osTypesFindByDescr", "fetching osTypes with description '%s'...", descr)

	p := a.client.OsType.NewListOsTypesParams()
	p.SetDescription(descr)
	resp, err := a.client.OsType.ListOsTypes(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not fetch osTypes with description '%s'", descr)
		return nil, err
	}

	a.logger.Debug("osTypesFindByDescr", "finished fetching osTypes with description '%s'", descr)
	return resp.OsTypes, nil
}

func (a CPI) osTypeFindByDescr(descr string) (*cloudstack.OsType, error) {
	osTypes, err := a.osTypesFindByDescr(descr)
	if err != nil {
		return nil, err
	}

	if len(osTypes) == 0 {
		err = bosherr.Errorf("could not find osType matching description '%s'", descr)
		return nil, err
	}

	if len(osTypes) > 1 {
		err = bosherr.Errorf("found multiple osTypes matching description '%s'", descr)
		return nil, err
	}

	return osTypes[0], nil
}
