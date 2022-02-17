package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
)

// CalculateVMCloudProperties
// find most-matching available disk and compute offers based on vm spec given in `vm_resources`
// key. https://bosh.io/docs/manifest-v2/#instance-groups
//
func (a CPI) CalculateVMCloudProperties(res apiv1.VMResources) (apiv1.VMCloudProps, error) {
	a.logger.Info("calculate_vm_cloud_properties",
		"finding offerings for disk:%d MB, ram:%d MB, cpu:%d...",
		res.EphemeralDiskSize, res.RAM, res.CPU,
	)

	diskOffering, err := a.diskOfferingFindBest(res.EphemeralDiskSize)
	if err != nil {
		res := apiv1.NewVMCloudPropsFromMap(map[string]interface{}{})
		err := bosherr.WrapErrorf(err, "could not find matching ephemeral disk offering")
		return res, err
	}

	serviceOffering, err := a.serviceOfferingFindBest(res.RAM, res.CPU)
	if err != nil {
		res := apiv1.NewVMCloudPropsFromMap(map[string]interface{}{})
		err := bosherr.WrapErrorf(err, "could not find matching ephemeral disk offering")
		return res, err
	}

	data := map[string]interface{}{
		"ephemeral_disk_offering": diskOffering.Name,
		"compute_offering":        serviceOffering.Name,
	}

	a.logger.Info("calculate_vm_cloud_properties", "finished finding offerings for disk:%d MB, ram:%d MB, cpu:%d (%#v)", data)
	return apiv1.NewVMCloudPropsFromMap(data), nil
}
