package reg

import "github.com/cloudfoundry/bosh-cpi-go/apiv1"

type AgentEnvService interface {
	// Fetch will return an error if Update was not called beforehand
	Fetch() (apiv1.AgentEnv, error)
	Delete() error
	Update(apiv1.AgentEnv) error
}
