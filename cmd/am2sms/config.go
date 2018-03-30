package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"

	"github.com/romana/rlog"
)

func CreateConfig() *appconfig.AppConfig {
	configFileName := flag.String("config", "", "Configuration file in YAML")
	flag.Parse()

	conf := appconfig.New(*configFileName)
	if conf == nil {
		_ = fmt.Errorf("Configuration error !")
		os.Exit(1)
	}
	rlog.Debug("Config =", conf)
	return conf
}
