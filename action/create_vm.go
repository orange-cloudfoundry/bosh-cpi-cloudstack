package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"fmt"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/satori/go.uuid"
)

func (a CPI) CreateVM(
	agentID apiv1.AgentID, stemcellCID apiv1.StemcellCID,
	cloudProps apiv1.VMCloudProps, networks apiv1.Networks,
	associatedDiskCIDs []apiv1.DiskCID, env apiv1.VMEnv) (apiv1.VMCID, error) {

	var resProps ResourceCloudProperties
	err := cloudProps.As(&resProps)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Cannot create vm")
	}

	diskCid, err := a.createEphemeralDisk(resProps.EphemeralDiskSize, resProps.DiskCloudProperties, nil)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapError(err, "Cannot create ephemeral disk when creating vm")
	}

	vmName := fmt.Sprintf("%s%s", config.VMPrefix, uuid.NewV4().String())

	userData := NewUserDataContents(vmName, a.config.Actions.Registry, networks)

	fact := apiv1.NewAgentEnvFactory()
	fact.ForVM()
	a.regFactory.Create("").Update()
	return apiv1.VMCID{}, nil
}
