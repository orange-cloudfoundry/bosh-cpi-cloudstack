package action

import (
	"fmt"
	"sort"
	"strings"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) diskOfferingsFindByName(offerName string) ([]*cloudstack.DiskOffering, error) {
	a.logger.Debug("diskOfferingsFindByName", "fetching disk offering with name '%s'...", offerName)

	p := a.client.DiskOffering.NewListDiskOfferingsParams()
	p.SetName(offerName)
	resp, err := a.client.DiskOffering.ListDiskOfferings(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not fetch disk offerings")
		a.logger.Error("diskOfferingsFindByName", err.Error())
		return nil, err
	}

	a.logger.Debug("diskOfferingsFindByName", "finish fetching disk offering with name '%s'...", offerName)
	return resp.DiskOfferings, nil
}

func (a CPI) diskOfferingFindByName(offerName string) (*cloudstack.DiskOffering, error) {
	offers, err := a.diskOfferingsFindByName(offerName)
	if err != nil {
		return nil, err
	}

	if len(offers) == 0 {
		err := bosherr.Errorf("could not find any disk offering with name '%s'", offerName)
		a.logger.Error("diskOfferingFindByName", err.Error())
		return nil, err
	}

	if len(offers) > 1 {
		err := bosherr.Errorf("found multiple disk offering with name '%s'", offerName)
		a.logger.Error("diskOfferingFindByName", err.Error())
		return nil, err

	}

	return offers[0], nil
}

// diskOfferingFindBest -
// find closest disk offering for given diskSizeMB
// 1. consider offers that match configuration tags and not-tags
// 2. sort by disk size asc
// 3. filter-out offers with not enough disk or custom
// 4. use offer with largest disk as fallback
//
func (a CPI) diskOfferingFindBest(diskSizeMB int) (*cloudstack.DiskOffering, error) {
	a.logger.Debug("diskOfferingFindBest", "fetching disk offering for size %d...", diskSizeMB)

	p := a.client.DiskOffering.NewListDiskOfferingsParams()
	resp, err := a.client.DiskOffering.ListDiskOfferings(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not fetch disk offerings")
		a.logger.Error("diskOfferingFindBest", err.Error())
		return nil, err
	}

	// 1.
	offers := a.diskOfferingFilterTags(resp.DiskOfferings)
	if len(offers) == 0 {
		tags := strings.Join(a.config.CloudStack.CalculateCloudProp.DiskTags, ",")
		err = fmt.Errorf("could not find disk offering matching tags '%s'", tags)
		a.logger.Error("diskOfferingFindBest", err.Error())
		return nil, err
	}

	// 2.
	sort.SliceStable(offers, func(i, j int) bool {
		return offers[i].Disksize < offers[j].Disksize
	})

	// 3.
	candidates := []*cloudstack.DiskOffering{}
	for _, cOffer := range offers {
		if (cOffer.Disksize * 1024) < int64(diskSizeMB) || cOffer.Iscustomized {
			continue
		}
		candidates = append(candidates, cOffer)
	}

	// 4.
	if len(candidates) == 0 {
		last := len(offers) - 1
		candidates = append(candidates, offers[last])
	}

	a.logger.Debug("diskOfferingFindBest",
		"finished fetching disk offering for size %d (%s/%s)...",
		diskSizeMB, candidates[0].Name, candidates[0].Disksize * 1024,
	)
	return candidates[0], nil
}


// diskOfferingFilterTags -
// filter available disk offering considering configuration
// - keep offers matchings DiskTags
// - reject offers matching NotDiskTags
//
func (a CPI) diskOfferingFilterTags(offers []*cloudstack.DiskOffering) []*cloudstack.DiskOffering {
	if len(a.config.CloudStack.CalculateCloudProp.DiskTags) == 0 &&
		len(a.config.CloudStack.CalculateCloudProp.NotDiskTags) == 0 {
		return offers
	}
	tmpOffers := make([]*cloudstack.DiskOffering, 0)
	for _, offer := range offers {
		for _, tag := range a.config.CloudStack.CalculateCloudProp.DiskTags {
			if strings.Contains(offer.Tags, tag) {
				tmpOffers = append(tmpOffers, offer)
				break
			}
		}
	}
	if len(tmpOffers) > 0 {
		offers = tmpOffers
		tmpOffers = make([]*cloudstack.DiskOffering, 0)
	}
	for _, offer := range offers {
		contains := false
		for _, tag := range a.config.CloudStack.CalculateCloudProp.NotDiskTags {
			if strings.Contains(offer.Tags, tag) {
				contains = true
				break
			}
		}
		if !contains {
			tmpOffers = append(tmpOffers, offer)
		}
	}
	return offers
}
