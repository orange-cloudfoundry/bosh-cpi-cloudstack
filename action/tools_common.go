package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func (a CPI) zoneFindByName(name string) (*cloudstack.Zone, error) {
	p := a.client.Zone.NewListZonesParams()
	p.SetName(name)
	resp, err := a.client.Zone.ListZones(p)
	if err != nil {
		return nil, err
	}
	if len(resp.Zones) == 0 {
		return nil, bosherr.Errorf("zone '%s' not found", name)
	}
	return resp.Zones[0], nil
}

func (a CPI) zoneFindDefault() (*cloudstack.Zone, error) {
	return a.zoneFindByName(a.config.CloudStack.DefaultZone)
}

func (a CPI) diskOfferingFindByName(name string) (*cloudstack.DiskOffering, error) {
	p := a.client.DiskOffering.NewListDiskOfferingsParams()
	p.SetName(name)

	resp, err := a.client.DiskOffering.ListDiskOfferings(p)
	if err != nil {
		return nil, err
	}
	if len(resp.DiskOfferings) == 0 {
		return nil, bosherr.Errorf("disk offering '%s' not found", name)
	}
	return resp.DiskOfferings[0], nil
}


func (a CPI) osTypeFindByDescr(descr string) (*cloudstack.OsType, error) {
	p := a.client.GuestOS.NewListOsTypesParams()
	p.SetDescription(descr)
	resp, err := a.client.GuestOS.ListOsTypes(p)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "could not list OSTypes with description '%s'", descr)
	}
	if resp.Count == 0 {
		return nil, bosherr.WrapErrorf(err, "could not find OSType with description '%s'", descr)
	}
	return resp.OsTypes[0], nil
}
