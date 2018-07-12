package fakes

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bwcvm "github.com/orange-cloudfoundry/bosh-cpi-cloudstack/vm"
)

type FakeMetadataService struct {
	Saved          bool
	SaveInstanceID apiv1.VMCID
	SaveErr        error
}

func NewFakeMetadataService() *FakeMetadataService {
	return &FakeMetadataService{}
}

func (ms *FakeMetadataService) Save(wardenFileService bwcvm.WardenFileService, instanceID apiv1.VMCID) error {
	ms.Saved = true
	ms.SaveInstanceID = instanceID
	return ms.SaveErr
}
