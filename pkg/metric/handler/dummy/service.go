package dummy

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"

	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

func (pm *dummyMetric) Config(conf *viper.Viper) error {
	rlog.Debug("(metric dummy).Config() called")
	return nil
}

func (pm *dummyMetric) IncCounter(name string) {
	rlog.Debug("(metric dummy).IncCounter(", name, ") called")
}

func (pm *dummyMetric) Info() domain.MetricHandlerInfo {
	return domain.MetricHandlerInfo{
		Name:    "dummy",
		Authors: "Stéphane Bunel",
		Version: "0.1",
	}
}
