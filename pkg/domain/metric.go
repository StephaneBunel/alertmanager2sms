package domain

import (
	"github.com/spf13/viper"
)

type (
	// Metricer is the metric interface
	Metricer interface {
		Config(conf *viper.Viper) error
		IncCounter(name string)
	}

	// MetricHandlerInfo should be passed when each metricer is registering itself
	MetricHandlerInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}
)
