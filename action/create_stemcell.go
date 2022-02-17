package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
)

func (a CPI) CreateStemcell(imagePath string, cp apiv1.StemcellCloudProps) (apiv1.StemcellCID, error) {
	name := a.templateGenerateID()
	a.logger.Info("create_stemcell", "creating stemcell '%s' for '%s'...", name, imagePath)

	err := a.createStemcell(name, imagePath, cp)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not create stemcell")
		a.logger.Error("create_stemcell", err.Error())
		return apiv1.StemcellCID{}, err
	}

	a.logger.Info("create_stemcell", "finished creating stemcell '%s' for '%s'", name, imagePath)
	return apiv1.NewStemcellCID(name), nil
}


// CreateStemcell - Create CS template from given stemcell
//
// 1. read cloud-properties
// 2. generate template name from random id
// 3. request CS an upload token
// 4. push image to CS recieved endpoint
func (a CPI) createStemcell(name string, imagePath string, cp apiv1.StemcellCloudProps) error {
	csProp := CloudStackCloudProperties{}
	err := cp.As(&csProp)
	if err != nil {
		return bosherr.WrapErrorf(err, "could not read cloud properties")
	}

	if err := csProp.Validate(); err != nil {
		return err
	}

	if len(csProp.LightTemplate) != 0 {
		return bosherr.Errorf("light stemcell is not supported")
	}

	zone, err := a.zoneFindDefault()
	if err != nil {
		return err
	}

	osType, err := a.osTypeFindByDescr(a.config.CloudStack.Stemcell.OsType)
	if err != nil {
		return err
	}

	uploadP, err := a.templateGetUploadParams(name, zone, osType)
	if err != nil {
		return err
	}

	r, err := a.templateUploadRequestCreate(imagePath, uploadP.PostURL, uploadP.Expires, uploadP.Signature, uploadP.Metadata)
	if err != nil {
		return err
	}

	err = a.templateUploadPerform(r)
	if err != nil {
		return err
	}

	err = a.templateInstallWait(uploadP.Id, name)
	if err != nil {
		return err
	}

	return nil
}
