package action

type DiskCloudProperties struct {
	DiskOffering          string `json:"disk_offering"`
	EphemeralDiskOffering string `json:"ephemeral_disk_offering"`
	EphemeralDiskSize     int    `json:"disk"`
}
