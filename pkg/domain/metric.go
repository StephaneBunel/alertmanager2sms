package domain

type (
	Metricer interface {
		Config()
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
