package action

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
)

// CreateStemcell - Create CS template from given stemcell
//
// 1. read cloud-properties
// 2. generate template name from random id
// 3. request CS an upload token
// 4. push image to CS received endpoint
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

	name := a.generateStemcellID()
	uploadP, err := a.getUploadParams(name)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell]")
	}

	a.logger.Info("create_stemcell", "Requesting upload to cloudstack for stemcell %s ...", name)
	r, err := a.createUploadRequest(imagePath, uploadP.PostURL, uploadP.Expires, uploadP.Signature, uploadP.Metadata)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell]")
	}
	a.logger.Info("create_stemcell", "Finished requesting upload to cloudstack for stemcell %s .", name)

	a.logger.Info("create_stemcell", "Uploading stemcell %s ...", name)
	err = a.performUpload(r)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell]")
	}

	err = a.pollTemplateStatus(uploadP.Id, name)
	if err != nil {
		return apiv1.StemcellCID{}, bosherr.WrapErrorf(err, "[create_stemcell]")
	}
	a.logger.Info("create_stemcell", "Finished Uploading stemcell %s ...", name)
	a.logger.Debug("create_stemcell", "create_stemcell success: template %s (%s)", uploadP.Id, name)
	return apiv1.NewStemcellCID(name), nil
}

// getUploadParams -
func (a CPI) getUploadParams(name string) (*cloudstack.GetUploadParamsForTemplateResponse, error) {
	zoneid, err := a.findZoneID()
	if err != nil {
		return nil, err
	}

	ostypeid, err := a.findOsTypeId(a.config.CloudStack.Stemcell.OsType)
	if err != nil {
		return nil, err
	}

	params := a.client.Template.NewGetUploadParamsForTemplateParams(
		name,
		config.TemplateFormat,
		config.Hypervisor,
		name,
		zoneid)
	params.SetOstypeid(ostypeid)
	params.SetIsextractable(true)
	params.SetRequireshvm(*a.config.CloudStack.Stemcell.RequiresHvm)
	params.SetBits(64)

	a.logger.Debug("create_stemcell", "requesting upload parameters : %#v", params)
	res, err := a.client.Template.GetUploadParamsForTemplate(params)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "[create_stemcell] could not get upload parameters")
	}

	return res, nil
}

// generateStemcellID -
// CS ids are limited to 32 bytes
func (a CPI) generateStemcellID() string {
	id := uuid.NewV4()
	name := fmt.Sprintf(config.TemplateNameFormat, id)
	parts := strings.Split(name, "-")
	return strings.Join(parts[0:4], "-")
}

// createUploadRequest -
func (a CPI) createUploadRequest(
	imagePath string,
	postURL string,
	expires string,
	signature string,
	metadata string) (*http.Request, error) {

	file, _ := os.Open(imagePath)
	defer file.Close()

	// header
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filepath.Base(file.Name())+".vhd.bz2"))
	h.Set("Content-Type", "application/x-bzip")

	// create part writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "unable to create multipart for image '%s'", imagePath)
	}

	// feed part writer
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "unable to read stemcell image '%s'", imagePath)
	}
	writer.Close()

	r, _ := http.NewRequest("POST", postURL, body)

	// manually replace header since:
	//  - CS is case-sensitive on header names
	//  - golang concert header names into UpperCamelCase
	r.Header = map[string][]string{
		"X-signature":  []string{signature},
		"X-metadata":   []string{metadata},
		"X-expires":    []string{expires},
		"Content-Type": []string{writer.FormDataContentType()},
	}

	return r, nil
}

// performUpload - upload stemcell to CS and check return status
// - InsecureSkipVerify since CSM is misconfigured
func (a CPI) performUpload(request *http.Request) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	a.logger.Debug("create_stemcell", "performing template upload to '%s'", request.URL)
	uploadRes, err := client.Do(request)
	if err != nil {
		return bosherr.WrapErrorf(err, "error while uploading stemcell")
	}

	if uploadRes.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(uploadRes.Body)
		return fmt.Errorf("unexpected response while uploading stemcell: %s", string(bodyBytes))
	}

	return nil
}

// pollTemplateStatus -
func (a CPI) pollTemplateStatus(templateID string, name string) error {
	dur, _ := time.ParseDuration(fmt.Sprintf("%ds", a.config.CloudStack.Timeout.PollTemplate))
	timeout := time.Now().Add(dur)
	for false == time.Now().After(timeout) {
		a.logger.Debug("create_stemcell", "checking status for template %s (%s)", templateID, name)
		resp, _, err := a.client.Template.GetTemplateByID(templateID, "executable")
		if err != nil {
			return bosherr.WrapErrorf(err, "unable to get status for template %s (%s)", templateID, name)
		}

		if resp.Status == "Download Complete" {
			return nil
		} else if strings.Contains(strings.ToLower(resp.Status), "error") {
			return fmt.Errorf("upload failed for template %s (%s) with status '%s'", templateID, name, resp.Status)
		}

		time.Sleep(5 * time.Second)
	}
	return fmt.Errorf("upload failed for template %s (%s), timeout reached", templateID, name)
}
