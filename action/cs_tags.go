package action

import (
	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func (a CPI) tagList(tagType string, resourceID string) ([]*cloudstack.Tag, error) {
	a.client.DefaultOptions()
	a.logger.Debug("tagList", "fetching tags for type '%s' and resource '%s'...", tagType, resourceID)

	p := a.client.Resourcetags.NewListTagsParams()
	p.SetResourcetype(tagType)
	p.SetResourceid(resourceID)
	resp, err := a.client.Resourcetags.ListTags(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not fetch tags for type '%s' and resource '%s'", tagType, resourceID)
		a.logger.Error("tagList", err.Error())
		return nil, err
	}

	a.logger.Debug("tagList", "finished fetching tags for type '%s' and resource '%s'", tagType, resourceID)
	return resp.Tags, nil
}

func (a CPI) tagDelete(tagType string, resourceID string) error {
	a.client.DefaultOptions()
	a.logger.Debug("tagDelete", "deleting tags for type '%s' and resource '%s'...", tagType, resourceID)

	p := a.client.Resourcetags.NewDeleteTagsParams([]string{resourceID}, tagType)
	_, err := a.client.Resourcetags.DeleteTags(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not delete tags for type '%s' and resource '%s'", tagType, resourceID)
		a.logger.Error("tagDelete", err.Error())
		return err
	}

	a.logger.Debug("tagDelete", "finished deleting tags for type '%s' and resource '%s'", tagType, resourceID)
	return nil
}

func (a CPI) tagCreate(tagType string, resourceID string, tags map[string]string) error {
	a.client.DefaultOptions()
	a.logger.Debug("tagCreate", "creating tags for type '%s' and resource '%s'...", tagType, resourceID)

	p := a.client.Resourcetags.NewCreateTagsParams([]string{resourceID}, string(tagType), tags)
	_, err := a.client.Resourcetags.CreateTags(p)
	if err != nil {
		err := bosherr.WrapErrorf(err, "could not create tags for type '%s' and resource '%s'", tagType, resourceID)
		a.logger.Error("tagCreate", err.Error())
		return err
	}

	a.logger.Debug("tagCreate", "finished creating tags for type '%s' and resource '%s'", tagType, resourceID)
	return nil
}
