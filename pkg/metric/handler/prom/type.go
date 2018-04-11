package prom

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
)

type (
	promMetric struct {
		promCounterMap map[string]prometheus.Counter
		once           sync.Once
		config         *viper.Viper
	}
)
