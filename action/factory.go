package action

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/xanzy/go-cloudstack/cloudstack"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/reg"
)

type Factory struct {
	config config.Config
	logger boshlog.Logger
}

type CPI struct {
	client     *cloudstack.CloudStackClient
	config     config.Config
	logger     boshlog.Logger
	regFactory reg.RegistryAgentFactory
}

func NewFactory(config config.Config, logger boshlog.Logger) Factory {
	return Factory{config, logger}
}

func (f Factory) New(_ apiv1.CallContext) (apiv1.CPI, error) {
	csConfig := f.config.CloudStack
	client := cloudstack.NewAsyncClient(csConfig.Endpoint, csConfig.ApiKey, csConfig.SecretAccessKey, csConfig.SkipVerifySSL)
	if f.config.CloudStack.Timeout.Global > 0 {
		client.AsyncTimeout(f.config.CloudStack.Timeout.Global)
	}

	regFactory := reg.NewFactory(f.config.Actions.Registry, f.logger)

	return &CPI{client, f.config, f.logger, regFactory}, nil
}
