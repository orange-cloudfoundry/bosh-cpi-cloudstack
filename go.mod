module github.com/orange-cloudfoundry/bosh-cpi-cloudstack

go 1.15

replace github.com/cppforlife/bosh-cpi-go => github.com/orange-cloudfoundry/bosh-cpi-go v1.1.7-ora

replace github.com/apache/cloudstack-go/v2 => github.com/orange-cloudfoundry/cloudstack-go/v2 v2.9.1-0.20210830121736-f98e41d9a2c3

require (
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/apache/cloudstack-go/v2 v2.0.0-00010101000000-000000000000
	github.com/cloudfoundry/bosh-utils v0.0.264
	github.com/cppforlife/bosh-cpi-go v0.0.0-20180718174221-526823bbeafd
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/prometheus/common v0.29.0
	github.com/prometheus/procfs v0.7.0 // indirect
	github.com/satori/go.uuid v1.2.0
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
