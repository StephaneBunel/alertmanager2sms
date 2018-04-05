package domain

type (
	Metricer interface {
		Config()
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
