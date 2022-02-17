package action

import (
	"time"
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

	a.logger.Info("periodic_clean_disk", "(background) cleaning emphemeral disks...")
	tags := map[string]string{
		"director": a.config.CloudStack.DirectorName,
	}
	volumes, err := a.volumesFindByTags(tags)
	if err != nil {
		a.logger.Warn("periodic_clean_disk", err.Error())
		return
	}

	volumes = a.volumesFilterEphemeral(volumes)
	volumes = a.volumesFilterDetached(volumes)
	volumes = a.volumesFilterReady(volumes)
	volumes = a.volumesFilterCreatedBefore(volumes, time.Now().Add(-1 * time.Hour))

	err = a.volumesDelete(volumes)
	if err != nil {
		a.logger.Warn("periodic_clean_disk", err.Error())
	}

	for _, cVolume := range volumes {
		a.logger.Info("periodic_clean_disk", "deleting volume %s (%s) ...", cVolume.Name, cVolume.Id)
		if err := a.volumeDelete(cVolume); err != nil {
			a.logger.Warn("periodic_clean_disk", "could not delete volume '%s' (%s)", cVolume.Name, cVolume.Id)
			continue
		}
	}
	a.logger.Info("periodic_clean_disk", "(background) finished cleaning emphemeral disks")
}
