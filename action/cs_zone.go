package action

import (
	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)


func (a CPI) zonesFindByName(name string) ([]*cloudstack.Zone, error) {
	a.logger.Debug("zonesFindByName", "fetching zones with name '%s'...", name)

	p := a.client.Zone.NewListZonesParams()
	p.SetName(name)
	resp, err := a.client.Zone.ListZones(p)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not fetch zones with name '%s'", name)
		return nil, err
	}

	a.logger.Debug("zonesFindByName", "finished fetching zones with name '%s'", name)
	return resp.Zones, nil
}

func (a CPI) zoneFindByName(name string) (*cloudstack.Zone, error) {
	zones, err := a.zonesFindByName(name)
	if err != nil {
		return nil, err
	}

	if len(zones) == 0 {
		err = bosherr.Errorf("could not find zone matching name '%s'", name)
		return nil, err
	}

	if len(zones) > 1 {
		err = bosherr.Errorf("found multiple zones matching name '%s'", name)
		return nil, err
	}

	return zones[0], nil
}

func (a CPI) zoneFindDefault() (*cloudstack.Zone, error) {
	return a.zoneFindByName(a.config.CloudStack.DefaultZone)
}
