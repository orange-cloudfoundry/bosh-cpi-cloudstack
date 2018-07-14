package main

import (
	"flag"
	"os"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cppforlife/bosh-cpi-go/rpc"

	bwcaction "github.com/orange-cloudfoundry/bosh-cpi-cloudstack/action"
	"github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

var (
	configPathOpt = flag.String("configPath", "", "Path to configuration file")
	cleaning      = flag.Bool("cleaning", false, "Set to true to run job for deleting periodically ephemeral disk")
)

func main() {
	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr, os.Stderr)
	defer logger.HandlePanic("Main")
	flag.Parse()

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
