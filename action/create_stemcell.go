package action

import (
	"bytes"
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// CreateStemcell - Create CS template from given stemcell
//
// 1. read cloud-properties
// 2. generate template name from random id
// 3. request CS an upload token
// 4. push image to CS recieved endpoint
func (a CPI) CreateStemcell(imagePath string, cp apiv1.StemcellCloudProps) (apiv1.StemcellCID, error) {
	csProp := CloudStackCloudProperties{}
	err := cp.As(&csProp)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell] error while reading cloud_properties")
	}
	if err = csProp.Validate(); err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell] unable to validate cloud_properties")
	}

	// TODO [xmt]: handle light stemcell properly
	if len(csProp.LightTemplate) != 0 {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell] not handling light stemcell yet")
	}

	name := fmt.Sprintf(config.TemplateNameFormat, uuid.Must(uuid.NewV4()))

	zoneid, err := a.findZoneId()
	if err != nil {
		return apiv1.StemcellCID{}, err
	}

	ostypeid, err := a.findOsTypeId(a.config.CloudStack.Stemcell.OsType)
	if err != nil {
		return apiv1.StemcellCID{}, err
	}

	// TODO [xmt]: check disk format
	params := a.client.Template.NewGetUploadParamsForTemplateParams(
		name,
		config.TemplateFormat,
		config.Hypervisor,
		name,
		ostypeid,
		zoneid)

	a.logger.Info("create_stemcell", "fds : %#v", params)

	res, err := a.client.Template.GetUploadParamsForTemplate(params)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell] could not get upload parameters")
	}
	a.logger.Info("create_stemcell", "fds : %#v", res)

	request, err := NewFileUploadRequest(res.PostURL, "file", imagePath)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell] could not prepare upload request for '%s'", imagePath)
	}

	request.Header.Set("X-signature", res.Signature)
	request.Header.Set("X-metadata", res.Metadata)
	request.Header.Set("X-expires", res.Expires)
	client := &http.Client{}
	if _, err = client.Do(request); err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell] error while uploading file '%s'", imagePath)
	}

	return apiv1.NewStemcellCID(res.Id), nil
}

// NewFileUploadRequest -
func NewFileUploadRequest(uri string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}
