package action

import (
	"strings"
	"time"

	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) PeriodicCleanDisk() {
	if a.config.CloudStack.DirectorName == "" {
		return
	}
	for {
		a.cleanDisk()
	}
}

func (a CPI) cleanDisk() {
	defer time.Sleep(time.Duration(a.config.CloudStack.IntervalCleanDisk) * time.Minute)
	a.logger.Info("periodic_clean_disk", "Start cleaning ephemeral disks ...")
	p := a.client.Volume.NewListVolumesParams()
	p.SetTags(map[string]string{
		"director": a.config.CloudStack.DirectorName,
	})

	resp, err := a.client.Volume.ListVolumes(p)
	if err != nil {
		a.logger.Warn("periodic_clean_disk", "Error occurred when finding volumes: %s", err.Error())
		return
	}

	for _, vol := range resp.Volumes {
		if !strings.HasPrefix(vol.Name, config.EphemeralDiskPrefix) ||
			vol.Vmname != "" ||
			vol.Destroyed {
			continue
		}

		t, err := time.Parse(time.RFC3339, vol.Created)
		if err != nil {
			a.logger.Warn("periodic_clean_disk", "Error occurred when parsing create time for volume %s: %s", vol.Name, err.Error())
			return
		}

		deleteTime := time.Now().Add(-1 * time.Hour)
		if t.Before(deleteTime) {
			return
		}
		a.logger.Info("periodic_clean_disk", "Deleting volume %s ...", vol.Name)
		delParams := a.client.Volume.NewDeleteVolumeParams(vol.Id)
		_, err = a.client.Volume.DeleteVolume(delParams)
		if err != nil {
			a.logger.Warn("periodic_clean_disk", "Error occurred when deleting volume %s: %s", vol.Name, err.Error())
			return
		}
		a.logger.Info("periodic_clean_disk", "Finished deleting volume %s .", vol.Name)
	}
	a.logger.Info("periodic_clean_disk", "Finished cleaning ephemeral disks.")

}
