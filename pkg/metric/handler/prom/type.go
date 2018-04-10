package prom

import (
	"sync"

	"github.com/spf13/viper"
)

type (
	promMetric struct {
		metrics map[string]interface{}
		once    sync.Once
		config  *viper.Viper
	}
)
