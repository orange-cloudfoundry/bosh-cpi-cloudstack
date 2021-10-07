package action

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

type UserDataService struct {
	logger boshlog.Logger
	isV2   bool
	data   UserData
}

type AgentSettings struct {
	AgentID   string              `json:"agent_id"`
	VM        apiv1.VMSpec        `json:"vm"`
	Mbus      string              `json:"mbus"`
	NTP       []string            `json:"ntp"`
	Blobstore apiv1.BlobstoreSpec `json:"blobstore"`
	Networks  apiv1.NetworksSpec  `json:"networks"`
	Disks     apiv1.DisksSpec     `json:"disks"`
	Env       apiv1.EnvSpec       `json:"env"`
}

type UserData struct {
	AgentSettings
	Registry UserDataRegistry `json:"registry"`
	Server   UserDataServer   `json:"server"`
	DNS      UserDataDNS      `json:"dns"`
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
	res.setRegistry(regOpts)
	res.setVM(vmName)
	return res
}

func (u *UserDataService) ToBase64() string {
	jsonStr, _ := json.Marshal(u.data)
	u.logger.Debug("UserDataService", "generated user-data: %s", jsonStr)
	base64Str := base64.StdEncoding.EncodeToString(jsonStr)
	return base64Str
}

func (u *UserDataService) SetAgentSettings(
	agentID apiv1.AgentID,
	cid apiv1.VMCID,
	networks apiv1.Networks,
	env apiv1.VMEnv,
	agentOptions apiv1.AgentOptions) {

	networksSpec := apiv1.NetworksSpec{}
	for netName, network := range networks {
		jsonStr, _ := json.Marshal(network)
		var spec apiv1.NetworkSpec
		_ = json.Unmarshal(jsonStr, &spec)
		networksSpec[netName] = spec
	}

	var envSpec apiv1.EnvSpec
	jsonStr, _ := json.Marshal(env)
	_ = json.Unmarshal(jsonStr, &envSpec)

	u.data.AgentSettings = AgentSettings{
		AgentID: agentID.AsString(),
		VM: apiv1.VMSpec{
			Name: cid.AsString(),
			ID:   cid.AsString(),
		},
		Mbus: agentOptions.Mbus,
		NTP:  agentOptions.NTP,
		Disks: apiv1.DisksSpec{
			System:    apiv1.NewDiskHintFromString("/dev/xvda"),
			Ephemeral: apiv1.NewDiskHintFromString("/dev/xvdb"),
		},
		Blobstore: apiv1.BlobstoreSpec{
			Provider: agentOptions.Blobstore.Type,
			Options:  agentOptions.Blobstore.Options,
		},
		Networks: networksSpec,
		Env:      envSpec,
	}
}

func (u *UserDataService) setVM(vmName string) {
	u.data.Server = UserDataServer{Name: vmName}
}

func (u *UserDataService) setDNS(networks apiv1.Networks) {
	list := make([]string, 0)
	defNet := networks.Default()
	if defNet.DNS() != nil {
		for _, dns := range defNet.DNS() {
			list = append(list, dns)
		}
	}
	u.data.DNS = UserDataDNS{Nameserver: list}
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

	u.data.Registry = UserDataRegistry{
		Endpoint: fmt.Sprintf("http://%s", host),
	}
}
