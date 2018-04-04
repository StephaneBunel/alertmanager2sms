package domain

type (
	IMetric interface {
		Config()
	}

	MetricInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}
)
