package dummy

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/romana/rlog"
)

func (pm *dummyMetric) Config() {
	rlog.Debug("(metric dummy).Config() called")
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
