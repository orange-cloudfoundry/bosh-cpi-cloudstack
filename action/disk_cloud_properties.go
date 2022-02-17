package action

import "github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"

type DiskCloudProperties struct {
	DiskOffering          string `json:"disk_offering"`
	EphemeralDiskOffering string `json:"ephemeral_disk_offering"`
	EphemeralDiskSize     int    `json:"disk"`
	RootDiskSize          int    `json:"root_disk_size"`
}

func (d DiskCloudProperties) GetDiskOffering(fallback config.DefaultOffer) string {
	if d.DiskOffering == "" {
		return fallback.Disk
	}
	return d.DiskOffering
}

func (d DiskCloudProperties) GetEphemeralDiskOffering(fallback config.DefaultOffer) string {
	if d.EphemeralDiskOffering == "" {
		return fallback.EphemeralDisk
	}
	return d.EphemeralDiskOffering
}
