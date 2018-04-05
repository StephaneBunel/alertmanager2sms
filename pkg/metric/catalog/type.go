package catalog

import "github.com/StephaneBunel/alertmanager2sms/pkg/domain"

type (
	MetricHandlerCataloger interface {
		Register(func() domain.Metricer, domain.MetricHandlerInfo)
		Exists(name string) bool
		New(name string) domain.Metricer
		ListByName() []string
	}

	metricCatalog map[string]func() domain.Metricer
)
