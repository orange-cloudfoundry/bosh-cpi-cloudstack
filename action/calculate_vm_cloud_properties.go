package action

import (
	"fmt"
	"sort"
	"strings"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) CalculateVMCloudProperties(res apiv1.VMResources) (apiv1.VMCloudProps, error) {

	ephemeralDiskOffering, err := a.findEphemeralDiskOffering(res.EphemeralDiskSize)
	if err != nil {
		return apiv1.NewVMCloudPropsFromMap(map[string]interface{}{}), bosherr.WrapErrorf(err, "Calculating vm cloud properties on ephemeral disk offerings")
	}
	serviceOffering, err := a.findServiceOffering(res.RAM, res.CPU)
	if err != nil {
		return apiv1.NewVMCloudPropsFromMap(map[string]interface{}{}), bosherr.WrapErrorf(err, "Calculating vm cloud properties on service offerings")
	}
	return apiv1.NewVMCloudPropsFromMap(map[string]interface{}{
		"compute_offering":        serviceOffering,
		"ephemeral_disk_offering": ephemeralDiskOffering,
	}), nil
}

// findEphemeralDiskOffering find a disk offering by disk size, expressed in megabytes
func (a CPI) findEphemeralDiskOffering(diskSize int) (string, error) {
	// convert to GiB
	diskSize /= 1024
	resp, err := a.client.DiskOffering.ListDiskOfferings(a.client.DiskOffering.NewListDiskOfferingsParams())
	if err != nil {
		return "", err
	}

	offers := a.filterDiskOffering(resp.DiskOfferings)
	if len(offers) == 0 {
		return "", fmt.Errorf("there is no offers corresponding to tags: %s", strings.Join(a.config.CloudStack.CalculateCloudProp.DiskTags, ","))
	}

	sort.SliceStable(offers, func(i, j int) bool {
		return offers[i].Disksize < offers[j].Disksize
	})

	finalOffers := make([]*cloudstack.DiskOffering, 0)
	for _, offer := range offers {
		if offer.Disksize < int64(diskSize) || offer.Iscustomized {
			continue
		}
		finalOffers = append(finalOffers, offer)
	}

	if len(finalOffers) == 0 {
		return offers[len(offers)-1].Name, nil
	}

	return finalOffers[0].Name, nil
}

func (a CPI) findServiceOffering(ram, cpu int) (string, error) {
	resp, err := a.client.ServiceOffering.ListServiceOfferings(a.client.ServiceOffering.NewListServiceOfferingsParams())
	if err != nil {
		return "", err
	}

	offers := resp.ServiceOfferings
	if len(offers) == 0 {
		return "", fmt.Errorf("there is no offers corresponding")
	}

	sort.SliceStable(offers, func(i, j int) bool {
		return offers[i].Memory < offers[j].Memory
	})

	ramOffers := make([]*cloudstack.ServiceOffering, 0)
	for _, offer := range offers {
		if offer.Memory < ram {
			continue
		}
		ramOffers = append(ramOffers, offer)
	}

	if len(ramOffers) == 0 {
		return offers[len(offers)-1].Name, nil
	}

	finalOffers := make([]*cloudstack.ServiceOffering, 0)
	for _, offer := range ramOffers {
		if offer.Cpunumber < cpu {
			continue
		}
		finalOffers = append(finalOffers, offer)
	}

	if len(finalOffers) == 0 {
		return ramOffers[len(ramOffers)-1].Name, nil
	}

	return finalOffers[0].Name, nil
}

func (a CPI) filterDiskOffering(offers []*cloudstack.DiskOffering) []*cloudstack.DiskOffering {
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
