package action

import (
	"fmt"
	"sort"
	"strings"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) serviceOfferingsGetAll() ([]*cloudstack.ServiceOffering, error) {
	a.logger.Debug("serviceOfferingGetAll", "fetching all compute offerings...")

	p := a.client.ServiceOffering.NewListServiceOfferingsParams()
	resp, err := a.client.ServiceOffering.ListServiceOfferings(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not fetch service offerings")
		a.logger.Error("serviceOfferingGetAll", err.Error())
		return nil, err
	}

	a.logger.Debug("serviceOfferingGetAll", "finished fetching all compute offerings...")
	return resp.ServiceOfferings, nil
}


func (a CPI) serviceOfferingsFindByName(name string) ([]*cloudstack.ServiceOffering, error) {
	a.logger.Debug("serviceOfferingsFindByName", "fetching compute offerings with name '%s'...", name)

	p := a.client.ServiceOffering.NewListServiceOfferingsParams()
	p.SetName(name)
	resp, err := a.client.ServiceOffering.ListServiceOfferings(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not fetch service offerings with name '%s'", name)
		a.logger.Error("serviceOfferingsFindByName", err.Error())
		return nil, err
	}

	a.logger.Debug("serviceOfferingsFindByName", "finished fetching compute offerings with name '%s'", name)
	return resp.ServiceOfferings, nil
}

func (a CPI) serviceOfferingFindByName(name string) (*cloudstack.ServiceOffering, error) {
	offers, err := a.serviceOfferingsFindByName(name)
	if err != nil {
		return nil, err
	}

	if len(offers) == 0 {
		err = bosherr.Errorf("could not find any service offering matching name '%s'", name)
		a.logger.Error("serviceOfferingFindByName", err.Error())
		return nil, err
	}

	if len(offers) > 1 {
		err = bosherr.Errorf("found multiple service offering matching name '%s'", name)
		a.logger.Error("serviceOfferingFindByName", err.Error())
		return nil, err
	}

	return offers[0], nil
}

// serviceOfferingFindBest -
// find closest compute offering for given ramMB and cpu count
// 1. consider offers that match configuration tags and not-tags
// 2. sort by service memory asc
// 3. filter-out offers with not enough memory
// 4. use offer with largest memory as fallback
// 5. sort by service cpu asc
// 6. filter-out offers with not enough cpu
// 7. use offer with largest cpu as fallback
//
func (a CPI) serviceOfferingFindBest(ramMB, cpu int) (*cloudstack.ServiceOffering, error) {
	a.logger.Debug("serviceOfferingFindBest", "fetching compute offering for ram:%d, cpu:%d...", ramMB, cpu)

	offers, err := a.serviceOfferingsGetAll()
	if err != nil {
		return nil, err
	}

	// 1.
	offers = a.serviceOfferingFilterTags(offers)
	if len(offers) == 0 {
		tags := strings.Join(a.config.CloudStack.CalculateCloudProp.ServiceTags, ",")
		err = fmt.Errorf("could not find service offering matching tags '%s'", tags)
		a.logger.Error("serviceOfferingFindBest", err.Error())
		return nil, err
	}

	// 2.
	sort.SliceStable(offers, func(i, j int) bool {
		return offers[i].Memory < offers[j].Memory
	})
	// 3.
	memoryCandidates := []*cloudstack.ServiceOffering{}
	for _, cOffer := range offers {
		if cOffer.Memory < ramMB {
			continue
		}
		memoryCandidates = append(memoryCandidates, cOffer)
	}
	// 4.
	if len(memoryCandidates) == 0 {
		last := len(offers) - 1
		memoryCandidates = append(memoryCandidates, offers[last])
	}

	// 5.
	sort.SliceStable(memoryCandidates, func(i, j int) bool {
		return memoryCandidates[i].Cpunumber < memoryCandidates[j].Cpunumber
	})
	// 6.
	candidates := []*cloudstack.ServiceOffering{}
	for _, cOffer := range memoryCandidates {
		if cOffer.Cpunumber < cpu {
			continue
		}
		candidates = append(candidates, cOffer)
	}
	// 7.
	if len(candidates) == 0 {
		last := len(memoryCandidates) - 1
		candidates = append(candidates, memoryCandidates[last])
	}

	a.logger.Debug("serviceOfferingFindBest",
		"finished fetching service offering for ram:%d, cpu:%d (%s/ram:%s/cpu:%s)...",
		ramMB, cpu, candidates[0].Name, candidates[0].Memory, candidates[0].Cpunumber,
	)
	return candidates[0], nil
}

func (a CPI) serviceOfferingFilterTags(offers []*cloudstack.ServiceOffering) []*cloudstack.ServiceOffering {
	if len(a.config.CloudStack.CalculateCloudProp.ServiceTags) == 0 &&
		len(a.config.CloudStack.CalculateCloudProp.NotServiceTags) == 0 {
		return offers
	}

	tmpOffers := make([]*cloudstack.ServiceOffering, 0)
	for _, offer := range offers {
		for _, tag := range a.config.CloudStack.CalculateCloudProp.ServiceTags {
			if strings.Contains(offer.Tags, tag) {
				tmpOffers = append(tmpOffers, offer)
				break
			}
		}
	}

	if len(tmpOffers) > 0 {
		offers = tmpOffers
		tmpOffers = make([]*cloudstack.ServiceOffering, 0)
	}

	for _, offer := range offers {
		contains := false
		for _, tag := range a.config.CloudStack.CalculateCloudProp.NotServiceTags {
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
