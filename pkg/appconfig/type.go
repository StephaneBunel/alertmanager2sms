package appconfig

import (
	"github.com/spf13/viper"
)

type (
	// AppConfig data structur
	AppConfig struct {
		Viper *viper.Viper
	}
)
