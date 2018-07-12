package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/xanzy/go-cloudstack/cloudstack"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Factory struct {
	config config.Config
	logger boshlog.Logger
}

type CPI struct {
	client *cloudstack.CloudStackClient
	config config.Config
	logger boshlog.Logger
}

func NewFactory(config config.Config, logger boshlog.Logger) Factory {
	return Factory{config, logger}
}

func (f Factory) New(_ apiv1.CallContext) (apiv1.CPI, error) {
	csConfig := f.config.CloudStack
	client := cloudstack.NewAsyncClient(csConfig.Endpoint, csConfig.ApiKey, csConfig.SecretAccessKey, csConfig.SkipVerifySSL)
	if f.config.CloudStack.AsyncTimeout > 0 {
		client.AsyncTimeout(f.config.CloudStack.AsyncTimeout)
	}

	return &CPI{client, f.config, f.logger}, nil
}
