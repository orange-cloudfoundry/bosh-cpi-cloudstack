module github.com/orange-cloudfoundry/bosh-cpi-cloudstack

go 1.15

replace github.com/cppforlife/bosh-cpi-go => github.com/orange-cloudfoundry/bosh-cpi-go v1.1.4-ora

require (
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/charlievieth/fs v0.0.1 // indirect
	github.com/cloudfoundry/bosh-utils v0.0.0-20210116100238-c508cff11a13
	github.com/cppforlife/bosh-cpi-go v0.0.0-20180718174221-526823bbeafd
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/orange-cloudfoundry/go-cloudstack v2.9.0-ora+incompatible
	github.com/prometheus/common v0.15.0
	github.com/satori/go.uuid v1.2.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
