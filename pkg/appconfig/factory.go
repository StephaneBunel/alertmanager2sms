package appconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

// New returns a new config object
func New(configFile string) *AppConfig {
	c := new(AppConfig)
	c.Viper = viper.New()

	if configFile == "" {
		c.Viper.SetConfigName("config")
		c.Viper.AddConfigPath(path.Join("/etc", c.AppName()))
		c.Viper.AddConfigPath(path.Join("$HOME", "."+c.AppName()))
		c.Viper.AddConfigPath(path.Join("$HOME", ".config", c.AppName()))
		c.Viper.AddConfigPath(path.Join("."))
		err := c.Viper.ReadInConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading configuration: %s", err)
			os.Exit(1)
		}
	} else {
		c.Viper.SetConfigType("yaml")
		fileData, err := ioutil.ReadFile(configFile)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		c.Viper.ReadConfig(bytes.NewBuffer(fileData))
	}

	cfKey := c.AppName() + ".logLevel"
	if c.Viper.IsSet(cfKey) {
		os.Setenv("RLOG_LOG_LEVEL", c.Viper.GetString(cfKey))
	}
	rlog.UpdateEnv()

	return c
}
