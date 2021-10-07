package reg

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

type RegistryAgentFactory struct {
	registryOptions config.RegistryOptions
	logger          boshlog.Logger
}

func NewFactory(registryOptions config.RegistryOptions, logger boshlog.Logger) RegistryAgentFactory {
	return RegistryAgentFactory{registryOptions, logger}
}

func (f RegistryAgentFactory) Create(instanceID apiv1.VMCID) AgentEnvService {
	if f.registryOptions.Host == "" {
		return NewNullAgentEnvService()
	}
	return NewRegistryAgentEnvService(f.registryOptions, instanceID, f.logger)
}
