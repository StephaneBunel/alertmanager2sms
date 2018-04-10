package domain

import (
	"github.com/spf13/viper"
)

type (
	Metricer interface {
		Config(conf *viper.Viper) error
		IncCounter(name string)
	}

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
