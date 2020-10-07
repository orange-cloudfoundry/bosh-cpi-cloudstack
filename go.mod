module github.com/orange-cloudfoundry/bosh-cpi-cloudstack

go 1.15

replace github.com/cppforlife/bosh-cpi-go => github.com/orange-cloudfoundry/bosh-cpi-go v1.1.3-ora

require (
	github.com/bmatcuk/doublestar v1.2.2 // indirect
	github.com/charlievieth/fs v0.0.0-20170613215519-7dc373669fa1 // indirect
	github.com/cloudfoundry/bosh-utils v0.0.0-20200926100154-fe1f0beb6a7d
	github.com/cppforlife/bosh-cpi-go v0.0.0-20180718174221-526823bbeafd
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/orange-cloudfoundry/go-cloudstack v0.0.0-20200212125423-ac26bbce93e3
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.6.1 // indirect
)
