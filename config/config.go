package config

import (
	"encoding/json"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"io/ioutil"
)

type RegistryOptions struct {
	Host     string
	Port     int
	Username string
	Password string
}

type FactoryOpts struct {
	Agent apiv1.AgentOptions

	AgentEnvService string
	Registry        RegistryOptions
}

func (o FactoryOpts) Validate() error {
	err := o.Agent.Validate()
	if err != nil {
		return bosherr.WrapError(err, "Validating Agent configuration")
	}

	return nil
}

type Config struct {
	CloudStack CloudStackConfig

	Actions FactoryOpts
}

type CloudStackConfig struct {
	Stemcell *StemcellConfig

	CalculateCloudProp CalculateCloudProp

	// Access
	Endpoint        string
	ApiKey          string
	SecretAccessKey string
	SkipVerifySSL   bool
	Timeout         Timeout

	DataDiskOffering string

	// Key
	DefaultKeyName string
	PrivateKey     string

	// Zone
	DefaultZone  string
	DefaultOffer DefaultOffer

	// VM
	ExpungeVm bool

	// PeriodicCleanDisk
	DirectorName      string
	IntervalCleanDisk int64

	EnableAutoAntiAffinity bool
}

type DefaultOffer struct {
	Disk          string
	EphemeralDisk string
	CustomDisk    string
}

type CalculateCloudProp struct {
	NotDiskTags    []string
	DiskTags       []string
	ServiceTags    []string
	NotServiceTags []string
}

type Timeout struct {
	Global               int64
	RebootVm             int64
	StopVm               int64
	CreateVm             int64
	DeleteVm             int64
	CreateVolume         int64
	DeleteVolume         int64
	ResizeVolume         int64
	PollTemplate         int64
	AttachVolume         int64
	DetachVolume         int64
	SnapshotVolume       int64
	DeleteSnapshotVolume int64
}

type StemcellConfig struct {
	PublicVisibility *bool
	RequiresHvm      *bool
	OsType           string
}

func NewConfigFromPath(path string) (Config, error) {
	config := defaultConfig()
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return config, bosherr.WrapErrorf(err, "Reading config '%s'", path)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, bosherr.WrapError(err, "Unmarshalling config")
	}

	err = config.Validate()
	if err != nil {
		return config, bosherr.WrapError(err, "Validating config")
	}

	return config, nil
}

func defaultConfig() Config {
	timeout := Timeout{
		Global:               1800,
		CreateVolume:         1800,
		DeleteVolume:         1800,
		RebootVm:             1800,
		CreateVm:             1800,
		DeleteVm:             1800,
		ResizeVolume:         1800,
		PollTemplate:         1800,
		AttachVolume:         36000,
		DetachVolume:         1800,
		StopVm:               1800,
		DeleteSnapshotVolume: 1800,
		SnapshotVolume:       1800,
	}
	defOffers := DefaultOffer{
		CustomDisk: "shared.custom",
	}
	return Config{
		CloudStack: CloudStackConfig{
			ExpungeVm:         true,
			Timeout:           timeout,
			DefaultOffer:      defOffers,
			IntervalCleanDisk: 60,
		},
	}
}

func (c Config) Validate() error {
	err := c.CloudStack.Validate()
	if err != nil {
		return bosherr.WrapError(err, "Validating CloudStack configuration")
	}

	err = c.Actions.Validate()
	if err != nil {
		return bosherr.WrapError(err, "Validating Actions configuration")
	}

	return nil
}

func (c CloudStackConfig) Validate() error {
	if c.Endpoint == "" {
		return bosherr.Error("Must provide non-empty Endpoint")
	}

	if c.ApiKey == "" {
		return bosherr.Error("Must provide non-empty ApiKey")
	}

	if c.SecretAccessKey == "" {
		return bosherr.Error("Must provide non-empty SecretAccessKey")
	}

	if c.DefaultZone == "" {
		return bosherr.Error("Must provide non-empty DefaultZone")
	}

	// TODO: maybe enforce use ssh keys

	err := c.Stemcell.Validate()
	if err != nil {
		return bosherr.WrapError(err, "Validating Stemcell configuration")
	}

	return nil
}

func (c *StemcellConfig) Validate() error {
	if c.PublicVisibility == nil {
		public := true
		c.PublicVisibility = &public
	}

	if c.RequiresHvm == nil {
		requires := true
		c.RequiresHvm = &requires
	}

	if c.OsType == "" {
		return bosherr.Error("Must provide non-empty OsType")
	}

	// TODO: maybe enforce use ssh keys

	return nil
}
