package action

import (
	"time"
	"strings"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

func (a CPI) PeriodicCleanDisk() {
	if a.config.CloudStack.DirectorName == "" {
		return
	}
	for {
		a.logger.Info("periodic_clean_disk", "Start cleaning emphemeral disks ...")
		p := a.client.Volume.NewListVolumesParams()
		p.SetTags(map[string]string{
			"director": a.config.CloudStack.DirectorName,
		})

		resp, err := a.client.Volume.ListVolumes(p)
		if err != nil {
			a.logger.Warn("periodic_clean_disk", "Error occured when finding volumes: %s", err.Error())
		}

		for _, vol := range resp.Volumes {
			if !strings.HasPrefix(vol.Name, config.EphemeralDiskPrefix) ||
				vol.Vmname != "" ||
				vol.Destroyed {
				continue
			}

			t, err := time.Parse(time.RFC3339, vol.Created)
			if err != nil {
				a.logger.Warn("periodic_clean_disk", "Error occured when parsing create time for volume %s: %s", vol.Name, err.Error())
				continue
			}

			deleteTime := time.Now().Add(-1 * time.Hour)
			if t.Before(deleteTime) {
				continue
			}
			a.logger.Info("periodic_clean_disk", "Deleting volume %s ...", vol.Name)
			delParams := a.client.Volume.NewDeleteVolumeParams(vol.Id)
			_, err = a.client.Volume.DeleteVolume(delParams)
			if err != nil {
				a.logger.Warn("periodic_clean_disk", "Error occured when deleting volume %s: %s", vol.Name, err.Error())
				continue
			}
			a.logger.Info("periodic_clean_disk", "Finished deleting volume %s .", vol.Name)
		}
		a.logger.Info("periodic_clean_disk", "Finished cleaning emphemeral disks.")
		time.Sleep(time.Duration(a.config.CloudStack.IntervalCleanDisk) * time.Minute)
	}
}
