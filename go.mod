module github.com/orange-cloudfoundry/bosh-cpi-cloudstack

go 1.14

replace github.com/cppforlife/bosh-cpi-go => github.com/orange-cloudfoundry/bosh-cpi-go v1.1.2-ora

require (
	github.com/bmatcuk/doublestar v1.2.2 // indirect
	github.com/charlievieth/fs v0.0.0-20170613215519-7dc373669fa1 // indirect
	github.com/cloudfoundry/bosh-utils v0.0.0-20200404100238-add7921dad2d
	github.com/cppforlife/bosh-cpi-go v0.0.0-20180718174221-526823bbeafd
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/orange-cloudfoundry/go-cloudstack v0.0.0-20200212125423-ac26bbce93e3
	github.com/satori/go.uuid v1.2.0
)
