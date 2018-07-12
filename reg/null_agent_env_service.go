package reg

import "github.com/cppforlife/bosh-cpi-go/apiv1"

type nullAgentEnvService struct {
}

func (nullAgentEnvService) Fetch() (apiv1.AgentEnv, error) {
	return &apiv1.AgentEnvImpl{}, nil
}

func (nullAgentEnvService) Update(apiv1.AgentEnv) error {
	return nil
}

func NewNullAgentEnvService() AgentEnvService {
	return &nullAgentEnvService{}
}
