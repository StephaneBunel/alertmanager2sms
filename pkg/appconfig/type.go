package appconfig

import (
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Viper *viper.Viper
	}
)
