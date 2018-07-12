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
	DefaultZone string

	DefaultOffers DefaultOffers
}

type DefaultOffers struct {
	Disk          string
	EphemeralDisk string
}

type CalculateCloudProps struct {
	DiskTags    []string
	ServiceTags []string
}

type Timeout struct {
	Global       int64
	Reboot       int64
	CreateVm     int64
	DeleteVm     int64
	CreateVolume int64
	DeleteVolume int64
}

func (t *Timeout) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if t.Global <= 0 {
		t.Global = 1800
	}
	if t.Reboot <= 0 {
		t.Reboot = t.Global
	}
	if t.CreateVm <= 0 {
		t.CreateVm = t.Global
	}
	if t.DeleteVm <= 0 {
		t.DeleteVm = t.Global
	}
	if t.CreateVolume <= 0 {
		t.CreateVolume = t.Global
	}
	if t.DeleteVolume <= 0 {
		t.DeleteVolume = t.Global
	}
	return nil
}

type StemcellConfig struct {
	PublicVisibility *bool
	RequiresHvm      *bool
	OsType           string
}

func NewConfigFromPath(path string, fs boshsys.FileSystem) (Config, error) {
	var config Config
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
