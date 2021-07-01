module github.com/orange-cloudfoundry/bosh-cpi-cloudstack

go 1.15

replace github.com/cppforlife/bosh-cpi-go => github.com/orange-cloudfoundry/bosh-cpi-go v1.1.5-ora

require (
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/cloudfoundry/bosh-utils v0.0.263
	github.com/cppforlife/bosh-cpi-go v0.0.0-20180718174221-526823bbeafd
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/orange-cloudfoundry/go-cloudstack v2.9.0-ora+incompatible
	github.com/prometheus/common v0.29.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
