package action

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	uuid "github.com/satori/go.uuid"
)

// templateGenerateID -
// CS ids are limited to 32 bytes
func (a CPI) templateGenerateID() string {
	id := uuid.NewV4()
	name := fmt.Sprintf(config.TemplateNameFormat, id)
	parts := strings.Split(name, "-")
	return strings.Join(parts[0:4], "-")
}



// templateGetUploadParams -
func (a CPI) templateGetUploadParams(name string, zone *cloudstack.Zone, osType *cloudstack.OsType) (*cloudstack.GetUploadParamsForTemplateResponse, error) {
	params := a.client.Template.NewGetUploadParamsForTemplateParams(
		name,
		config.TemplateFormat,
		config.Hypervisor,
		name,
		zone.Id)
	params.SetOstypeid(osType.Id)
	params.SetIsextractable(true)
	params.SetRequireshvm(*a.config.CloudStack.Stemcell.RequiresHvm)
	params.SetBits(64)

	a.logger.Debug("templateGetUploadParams", "requesting upload parameters for template '%s'...", name)
	a.logger.Debug("templateGetUploadParams", "params: %#v", params)
	res, err := a.client.Template.GetUploadParamsForTemplate(params)
	if err != nil {
		err = bosherr.WrapErrorf(err, "could not get template upload parameters")
		a.logger.Error("templateGetUploadParams", err.Error())
		return nil, err
	}

	a.logger.Debug("templateGetUploadParams", "finished requesting upload parameters for template '%s'", name)
	return res, nil
}


// templateUploadRequestCreate - prepare an http request that upload template on endpoint given by
// previous GetUploadParamsForTemplate call
//
// 1. preapre headers
// 2. create multipart writer
// 3. feed multipart writer
// 4. manually replace header since:
//    - CS is case sensitive on header names
//    - golang concert header names into UpperCamelCase
func (a CPI) templateUploadRequestCreate(
	name string,
	imagePath string,
	postURL string,
	expires string,
	signature string,
	metadata string) (*http.Request, error) {

	a.logger.Debug("templateUploadRequestCreate", "preparing template upload request for stemcell '%s' ...", name)
	a.logger.Debug("templateUploadRequestCreate", "path=%s, postURL=%s", imagePath, postURL)
	file, _ := os.Open(imagePath)
	defer file.Close()

	// 1.
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filepath.Base(file.Name())+".vhd.bz2"))
	h.Set("Content-Type", "application/x-bzip")

	// 2.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreatePart(h)
	if err != nil {
		err = bosherr.WrapErrorf(err, "unable to create multipart for image '%s'", imagePath)
		a.logger.Error("templateUploadRequestCreate", err.Error())
		return nil, err
	}

	// 3.
	_, err = io.Copy(part, file)
	if err != nil {
		err = bosherr.WrapErrorf(err, "unable to read stemcell image '%s'", imagePath)
		a.logger.Error("templateUploadRequestCreate", err.Error())
		return nil, err
	}
	writer.Close()

	r, _ := http.NewRequest("POST", postURL, body)

	// 4.
	r.Header = map[string][]string{
		"X-signature":  []string{signature},
		"X-metadata":   []string{metadata},
		"X-expires":    []string{expires},
		"Content-Type": []string{writer.FormDataContentType()},
	}

	a.logger.Debug("templateUploadRequestCreate", "finished preparing template upload request for stemcell '%s'", name)
	return r, nil
}

// templateUploadPerform - upload stemcell to CS and check return status
//
// 1. InsecureSkipVerify since CSM is misconfigured
func (a CPI) templateUploadPerform(request *http.Request) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// 1.
				InsecureSkipVerify: true,
			},
		},
	}

	a.logger.Debug("templateUploadPerform", "performing template upload to '%s'...", request.URL)
	uploadRes, err := client.Do(request)
	if err != nil {
		err = bosherr.WrapErrorf(err, "error while uploading stemcell")
		a.logger.Error("templateUploadPerform", err.Error())
		return err
	}

	if uploadRes.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(uploadRes.Body)
		err = fmt.Errorf("unexpected response while uploading stemcell: %s", string(bodyBytes))
		a.logger.Error("templateUploadPerform", err.Error())
		return err
	}

	a.logger.Debug("templateUploadPerform", "finished performing template upload to '%s'", request.URL)
	return nil

}


// templateUploadWait - wait until cloudstack has successfully installed uploaded template into
// its system or until timeout is reached.
func (a CPI) templateInstallWait(templateID string, name string) error {
	a.logger.Debug("templateInstallWait", "checking install status for template '%s' (%s)...", templateID, name)
	dur, _ := time.ParseDuration(fmt.Sprintf("%ds", a.config.CloudStack.Timeout.PollTemplate))
	timeout := time.Now().Add(dur)

	for false == time.Now().After(timeout) {
		resp, _, err := a.client.Template.GetTemplateByID(templateID, "executable")
		if err != nil {
			err = bosherr.WrapErrorf(err, "unable to get status for template %s (%s)", templateID, name)
			a.logger.Error("templateInstallWait", err.Error())
			return err
		}

		if resp.Status == "Download Complete" {
			a.logger.Debug("templateInstallWait", "finished checking install status for template '%s' (%s): template is ready", templateID, name)
			return nil
		}

		if strings.Contains(strings.ToLower(resp.Status), "error") {
			err = fmt.Errorf("upload failed for template %s (%s) with status '%s'", templateID, name, resp.Status)
			a.logger.Error("templateInstallWait", err.Error())
			return err
		}

		a.logger.Debug("templateInstallWait", "checking install status for template '%s' (%s): template is not ready retrying...", templateID, name)
		time.Sleep(5 * time.Second)
	}

	err := fmt.Errorf("upload failed for template %s (%s), timeout reached", templateID, name)
	a.logger.Error("templateInstallWait", err.Error())
	return err
}

func (a CPI) templateFindByName(name string, zone *cloudstack.Zone) (*cloudstack.Template, error) {
	a.logger.Debug("templateFindByName", "fetching template '%s'...", name)
	template, _, err := a.client.Template.GetTemplateByName(name, "executable", zone.Id)
	if err != nil {
		err = bosherr.WrapErrorf(err, "unable to find template name '%s'", name)
		a.logger.Error("templateFindByName", err.Error())
		return nil, err
	}
	a.logger.Debug("templateFindByName", "finished fetching template '%s' (%d)", name, template.Id)
	return template, nil
}


func (a CPI) templateDelete(template *cloudstack.Template, zone *cloudstack.Zone) error {
	a.client.AsyncTimeout(a.config.CloudStack.Timeout.DeleteStemcell)

	params := a.client.Template.NewDeleteTemplateParams(template.Id)
	params.SetZoneid(zone.Id)

	a.logger.Debug("templateDelete", "deleting template '%s' (%s)...", template.Id, template.Name)
	_, err := a.client.Template.DeleteTemplate(params)
	if err != nil {
		err = bosherr.WrapErrorf(err, "unable to delete template name '%s' (%s)", template.Id, template.Name)
		a.logger.Error("templateDelete", err.Error())
		return err
	}
	a.logger.Debug("templateDelete", "finished deleting template '%s' (%s)...", template.Id, template.Name)
	return nil
}
