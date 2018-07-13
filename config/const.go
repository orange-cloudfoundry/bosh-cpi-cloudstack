package config

import "strings"

type Tags string

type DiskType string

type NetworkType string

type VmState string

const (
	UserVm                Tags        = "UserVm"
	Volume                Tags        = "Volume"
	Snapshot              Tags        = "Snapshot"
	Root                  DiskType    = "ROOT"
	Datadisk              DiskType    = "DATADISK"
	ManualNetwork         NetworkType = "manual"
	DynamicNetwork        NetworkType = "dynamic"
	VipNetwork            NetworkType = "vip"
	VmRunning             VmState     = "running"
	PersistenceDiskPrefix             = "cpi-disk-"
	EphemeralDiskPrefix               = "cpi-ephemeral-disk-"
	VMPrefix                          = "cpivm-"
	TemplateFormat                    = "VHD"
	TemplateNameFormat                = "cpitemplate-%s"
	Hypervisor                        = "XenServer"
)

func ToVmState(state string) VmState {
	return VmState(strings.ToLower(state))
}
