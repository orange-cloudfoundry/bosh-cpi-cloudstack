module github.com/orange-cloudfoundry/bosh-cpi-cloudstack

go 1.15

replace github.com/cppforlife/bosh-cpi-go => github.com/orange-cloudfoundry/bosh-cpi-go v1.1.5-ora

require (
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/cloudfoundry/bosh-utils v0.0.0-20210320100230-b112c198f4b7
	github.com/cppforlife/bosh-cpi-go v0.0.0-20180718174221-526823bbeafd
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/onsi/ginkgo v1.15.2
	github.com/onsi/gomega v1.11.0
	github.com/orange-cloudfoundry/go-cloudstack v2.9.0-ora+incompatible
	github.com/prometheus/client_golang v1.10.0 // indirect
	github.com/prometheus/common v0.20.0
	github.com/satori/go.uuid v1.2.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
