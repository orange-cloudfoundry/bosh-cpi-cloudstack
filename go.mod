module github.com/orange-cloudfoundry/bosh-cpi-cloudstack

go 1.15

replace github.com/cppforlife/bosh-cpi-go => github.com/orange-cloudfoundry/bosh-cpi-go v1.1.7-ora

require (
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/apache/cloudstack-go/v2 v2.9.1-0.20210831123014-8290b0373f69
	github.com/cloudfoundry/bosh-utils v0.0.273
	github.com/cppforlife/bosh-cpi-go v0.0.0-20180718174221-526823bbeafd
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/prometheus/common v0.29.0
	github.com/prometheus/procfs v0.7.0 // indirect
	github.com/satori/go.uuid v1.2.0
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
