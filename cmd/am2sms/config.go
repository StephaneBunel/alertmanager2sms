package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"

	"github.com/romana/rlog"
)

var (
	VersionBuild = "N/A" // Passed at compile time
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s (build: %s)\n", os.Args[0], VersionBuild)
	flag.PrintDefaults()
}

// CreateConfig returns a new appconfig object
func CreateConfig() *appconfig.AppConfig {
	configFileName := flag.String("config", "", "Configuration file in YAML")
	flag.Usage = usage
	flag.Parse()

	conf := appconfig.New(*configFileName)
	if conf == nil {
		_ = fmt.Errorf("configuration error")
		os.Exit(1)
	}
	rlog.Debug("Config =", conf)
	rlog.Info("config successfully created")
	return conf
}
