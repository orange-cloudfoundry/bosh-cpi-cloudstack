package action

import (
	"fmt"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/reg"
)

type Factory struct {
	config config.Config
	logger boshlog.Logger
}

type Context struct {
	DirectorUUID string `json:"director_uuid"`
	APIVersion   int    `json:"api_version"`
}

type CPI struct {
	client     *cloudstack.CloudStackClient
	config     config.Config
	logger     boshlog.Logger
	regFactory reg.RegistryAgentFactory
	ctx        Context
}

func NewFactory(config config.Config, logger boshlog.Logger) Factory {
	return Factory{config, logger}
}

func (f Factory) New(callCtx apiv1.CallContext) (apiv1.CPI, error) {
	csConfig := f.config.CloudStack
	client := cloudstack.NewAsyncClient(csConfig.Endpoint, csConfig.ApiKey, csConfig.SecretAccessKey, csConfig.SkipVerifySSL)
	if f.config.CloudStack.Timeout.Global > 0 {
		client.AsyncTimeout(f.config.CloudStack.Timeout.Global)
	}

	regFactory := reg.NewFactory(f.config.Actions.Registry, f.logger)

	var ctx Context
	if callCtx != nil {
		err := callCtx.As(&ctx)
		if err != nil {
			return nil, fmt.Errorf("error while reading call context: %v", err)
		}
	}

	return &CPI{client, f.config, f.logger, regFactory, ctx}, nil
}
