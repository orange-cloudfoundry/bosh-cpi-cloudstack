package main

import (
	"os"

	"github.com/cloudfoundry/bosh-cpi-go/rpc"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/alecthomas/kingpin/v2"
	bwcaction "github.com/orange-cloudfoundry/bosh-cpi-cloudstack/action"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
	"github.com/prometheus/common/version"
)

var (
	configPathOpt = kingpin.Flag("configPath", "Configuration file path").Required().String()
	cleaning      = kingpin.Flag("cleaning", "Set to true to run job for deleting periodically ephemeral disk").Bool()
)

func main() {
	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr)
	defer logger.HandlePanic("Main")

	kingpin.Version(version.Print("bosh-cpi-cloudstack"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	c, err := config.NewConfigFromPath(*configPathOpt)
	if err != nil {
		logger.Error("main", "Loading config %s", err.Error())
		os.Exit(1)
	}

	cpiFactory := bwcaction.NewFactory(c, logger)

	if *cleaning {
		cpi, _ := cpiFactory.New(nil)
		cpi.(*bwcaction.CPI).PeriodicCleanDisk()
		return
	}

	cli := rpc.NewFactory(logger).NewCLI(cpiFactory)

	err = cli.ServeOnce()
	if err != nil {
		logger.Error("main", "Serving once %s", err)
		os.Exit(1)
	}
}
