package action

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

type UserDataService struct {
	logger boshlog.Logger
	isV2   bool
	data   UserData
}

type UserData struct {
	Registry *UserDataRegistry `json:"registry,omitempty"`
	Server   *UserDataServer   `json:"server,omitempty"`
	DNS      *UserDataDNS      `json:"dns,omitempty"`
	Networks apiv1.Networks    `json:"networks,omitempty"`
	Disks    *apiv1.DisksSpec  `json:"disks,omitempty"`
	VM       *UserDataVM       `json:"vm,omitempty"`
}

type UserDataRegistry struct {
	Endpoint string `json:"endpoint"`
}

type UserDataServer struct {
	Name string `json:"name"`
}

type UserDataDNS struct {
	Nameserver []string `json:"nameserver"`
}

type UserDataVM struct {
	Name string `json:"name"`
}

func NewUserDataService(
	logger boshlog.Logger,
	vmName string,
	regOpts config.RegistryOptions,
	networks apiv1.Networks,
	isV2 bool,
) *UserDataService {

	res := &UserDataService{
		logger: logger,
		isV2:   isV2,
		data:   UserData{},
	}

	res.setDNS(networks)
	if isV2 {
		res.setDisks()
		res.setNetworks(networks)
	} else {
		res.setRegistry(regOpts)
	}
	res.setVM(vmName)
	return res
}

func (u *UserDataService) ToBase64() string {
	jsonStr, _ := json.Marshal(u.data)
	u.logger.Debug("UserDataService", "generated user-data: %s", jsonStr)
	base64Str := base64.StdEncoding.EncodeToString(jsonStr)
	return base64Str
}

func (u *UserDataService) setVM(vmName string) {
	if u.isV2 {
		u.data.VM = &UserDataVM{Name: vmName}
	} else {
		u.data.Server = &UserDataServer{Name: vmName}
	}
}

func (u *UserDataService) setNetworks(networks apiv1.Networks) {
	u.data.Networks = apiv1.Networks{}
	for name, network := range networks {
		if network.Type() != string(config.ManualNetwork) && network.Type() != string(config.DynamicNetwork) {
			continue
		}
		u.data.Networks[name] = network
	}
}

func (u *UserDataService) setDisks() {
	u.data.Disks = &apiv1.DisksSpec{
		System:    apiv1.NewDiskHintFromMap(map[string]interface{}{"path": "/dev/xvda"}),
		Ephemeral: apiv1.NewDiskHintFromMap(map[string]interface{}{"path": "/dev/xvdb"}),
	}
}

func (u *UserDataService) setDNS(networks apiv1.Networks) {
	list := make([]string, 0)
	defNet := networks.Default()
	if defNet.DNS() != nil {
		for _, dns := range defNet.DNS() {
			list = append(list, dns)
		}
	}
	u.data.DNS = &UserDataDNS{Nameserver: list}
}

func (u *UserDataService) setRegistry(opts config.RegistryOptions) {
	host := opts.Host
	if opts.Port != 0 {
		host = fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	}

	if opts.Username != "" {
		userInfo := opts.Username
		if opts.Password != "" {
			userInfo += fmt.Sprintf(":%s", opts.Password)
		}
		host = fmt.Sprintf("%s@%s", userInfo, host)
	}

	u.data.Registry = &UserDataRegistry{
		Endpoint: fmt.Sprintf("http://%s", host),
	}
}
