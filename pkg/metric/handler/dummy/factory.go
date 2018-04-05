package dummy

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/metric/catalog"
)

func New() domain.Metricer {
	h := new(dummyMetric)
	return h
}

func init() {
	catalog.Default().Register(New, (&dummyMetric{}).Info())
}
