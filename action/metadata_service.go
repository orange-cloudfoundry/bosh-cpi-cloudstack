package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"fmt"
)

type UserDataContentsType struct {
	Registry UserDataRegistry
	Server   UserDataServer
	DNS      UserDataDNS
}

type UserDataRegistry struct {
	Endpoint string
}

type UserDataServer struct {
	Name string // Name given by CPI e.g. vm-384sd4-r7re9e...
}

type UserDataDNS struct {
	Nameserver []string
}

func NewUserDataContents(vmName string, regOpts config.RegistryOptions, networks apiv1.Networks) UserDataContentsType {
	regUser := ""
	if regOpts.Username != "" {
		regUser += regOpts.Username
	}
	if regOpts.Username != "" && regOpts.Password != "" {
		regUser += ":" + regOpts.Password
	}
	if regUser != "" {
		regUser += "@"
	}
	regEndpoint := fmt.Sprintf("http://%s%s", regUser, regOpts.Host)
	if regOpts.Port != 0 {
		regEndpoint = fmt.Sprintf("%s:%d", regEndpoint, regOpts.Port)
	}

	dnsServer := make([]string, 0)
	defNet := networks.Default()

	if defNet.DNS() != nil && len(defNet.DNS()) > 0 {
		dnsServer = append(dnsServer, defNet.DNS()[0])
	}

	return UserDataContentsType{
		Registry: UserDataRegistry{regEndpoint},
		Server:   UserDataServer{vmName},
		DNS:      UserDataDNS{dnsServer},
	}
}
