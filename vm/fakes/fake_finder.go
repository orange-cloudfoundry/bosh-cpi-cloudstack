package fakes

import (
	bwcvm "github.com/orange-cloudfoundry/bosh-cpi-cloudstack/vm"
)

type FakeFinder struct {
	FindID    string
	FindVM    bwcvm.VM
	FindFound bool
	FindErr   error
}

func (f *FakeFinder) Find(id string) (bwcvm.VM, bool, error) {
	f.FindID = id
	return f.FindVM, f.FindFound, f.FindErr
}
