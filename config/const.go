package config

type Tags string

type DiskType string

const (
	UserVm                Tags     = "UserVm"
	Volume                Tags     = "Volume"
	Root                  DiskType = "ROOT"
	Datadisk              DiskType = "DATADISK"
	PersistenceDiskPrefix          = "cpi-disk-"
	EphemeralDiskPrefix            = "cpi-ephemeral-disk-"
	VMPrefix                       = "cpivm-"
	TemplateFormat                 = "VHD"
)
