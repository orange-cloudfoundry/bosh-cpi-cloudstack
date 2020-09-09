package action

import (
	"fmt"
)

type CloudStackCloudProperties struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	Infrastructure  string `json:"infrastructure"`
	Hypervisor      string `json:"hypervisor"`
	Disk            int    `json:"disk"`
	DiskFormat      string `json:"disk_format"`
	ContainerFormat string `json:"container_format"`
	OsType          string `json:"os_type"`
	OsDistro        string `json:"os_distro"`
	Architecture    string `json:"architecture"`
	AutoDiskConfig  bool   `json:"auto_disk_config"`
	LightTemplate   string `json:"light_template"`
}

func (cc CloudStackCloudProperties) Validate() error {
	if cc.Infrastructure != "cloudstack" {
		return fmt.Errorf("infrastructure '%s' is not supported (must be cloudstack)", cc.Infrastructure)
	}
	if cc.Architecture != "x86_64" {
		return fmt.Errorf("architecture '%s' is not supported (must be x86_64)", cc.Architecture)
	}
	if cc.Hypervisor != "xen" {
		return fmt.Errorf("hypervisor '%s' is not supported (must be xen)", cc.Architecture)
	}
	if cc.OsType != "linux" {
		return fmt.Errorf("os_type '%s' is not supported (must be linux)", cc.OsType)
	}
	return nil
}

type LBCloudProperties struct {
	Lbs []LBConfig `json:"lbs"`
}

type LBConfig struct {
	Name         string `json:"name"`
	Algorithm    string `json:"algorithm"`
	PrivatePort  int    `json:"private_port"`
	PublicPort   int    `json:"public_port"`
	OpenFirewall bool   `json:"open_firewall"`
	PublicIp     string `json:"public_ip"`
}

type ComputeCloudProperties struct {
	CPUNumber int `json:"cpu"`
	CPUSpeed  int `json:"cpu_speed"`
	RAM       int `json:"ram"`
}

type Route struct {
	Network     string
	Destination string
	Gateway     string
	Netmask     string
}

type VMExtRoute struct {
	CIDR    string
	Gateway string
}
type VMExtRoutes []VMExtRoute
type VMExtRouteMap map[string]VMExtRoutes

type ResourceCloudProperties struct {
	DiskCloudProperties
	LBCloudProperties
	ComputeCloudProperties
	Routes            VMExtRouteMap `json:"routes"`
	ComputeOffering   string        `json:"compute_offering"`
	AffinityGroup     string        `json:"affinity_group"`
	AffinityGroupType string        `json:"affinity_group_type"`
}

type NetworkCloudProperties struct {
	Name           string `json:"name"`
	UnDiscoverable bool   `json:"undiscoverable"`
}
