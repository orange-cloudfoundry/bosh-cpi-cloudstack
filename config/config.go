package config

import (
	"encoding/json"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

/*

   @Value("${cloudstack.state_timeout}")
   public int state_timeout;

   @Value("${cloudstack.state_timeout_volume}")
   public int state_timeout_volume;

   @Value("${cloudstack.stemcell_publish_timeout}")
   public int publishTemplateTimeoutMinutes;



   @Value("${cloudstack.default_zone}")
   public String default_zone;


   @Value("${cpi.vm_create_delay}")
   public int vmCreateDelaySeconds;


   @Value("${cpi.vm_expunge_delay}")
   public int vmExpungeDelaySeconds;

   @Value("${cpi.force_expunge}")
   public boolean forceVmExpunge;


   @Value("${cpi.default_disk_offering}")
   public String defaultDiskOffering;

   @Value("${cpi.default_ephemeral_disk_offering}")
   public String defaultEphemeralDiskOffering;

   @Value("${cpi.lightstemcell.instance_type}")
   public String light_stemcell_instance_type;//"CO1 - Small STD";

   @Value("${cpi.lightstemcell.network_name}")
   public String lightStemcellNetworkName; //"3112 - preprod - back";

   public List<String> calculateDiskTags;
   public List<String> calculateComputeTags;

   @Value("${cpi.calculate_vm_cloud_properties.disk.tags:#{null}}")
   private String calculateDiskTagsRaw;

   @Value("${cpi.calculate_vm_cloud_properties.compute.tags:#{null}}")
   private String calculateComputeTagsRaw;
*/

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

	CalculateCloudProps CalculateCloudProps

	// Access
	Endpoint        string
	ApiKey          string
	SecretAccessKey string
	SkipVerifySSL   bool
	Timeout         Timeout

	// Key
	DefaultKeyName string
	PrivateKey     string

	// Zone
	DefaultZone  string
	DefaultOffer DefaultOffer
}

type DefaultOffer struct {
	Disk          string
	EphemeralDisk string
	CustomDisk    string
}

type CalculateCloudProps struct {
	DiskTags    []string
	ServiceTags []string
}

type Timeout struct {
	Global       int64
	RebootVm     int64
	StopVm       int64
	CreateVm     int64
	DeleteVm     int64
	CreateVolume int64
	DeleteVolume int64
	ResizeVolume int64
	PollTemplate int64
	AttachVolume int64
	DetachVolume int64
}

type StemcellConfig struct {
	PublicVisibility *bool
	RequiresHvm      *bool
	OsType           string
}

func NewConfigFromPath(path string, fs boshsys.FileSystem) (Config, error) {
	config := defaultConfig()
	bytes, err := fs.ReadFile(path)
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
		Global:       1800,
		CreateVolume: 1800,
		DeleteVolume: 1800,
		RebootVm:     1800,
		CreateVm:     1800,
		DeleteVm:     1800,
		ResizeVolume: 1800,
		PollTemplate: 1800,
		AttachVolume: 1800,
		DetachVolume: 1800,
		StopVm:       1800,
	}
	defOffers := DefaultOffer{
		CustomDisk: "shared.custom",
	}
	return Config{
		CloudStack: CloudStackConfig{
			Timeout:      timeout,
			DefaultOffer: defOffers,
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
