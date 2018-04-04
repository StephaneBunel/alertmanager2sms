package metrics

/*
import (
	"strings"

	"github.com/romana/rlog"
)

func MetricCatalog() IMetricCatalog {
	return metricPC
}

func (c *metricProviderCatalog) Add(factory func() IMetric, info MetricInfo) {
	c.catalog[strings.ToLower(info.Name)] = factory
}

func (c *metricProviderCatalog) Exists(name string) bool {
	rlog.Debug(c)
	_, exists := c.catalog[strings.ToLower(name)]
	return exists
}

func (c *metricProviderCatalog) NewByName(name string) IMetric {
	if c.Exists(name) {
		return c.catalog[strings.ToLower(name)]()
	}
	return nil
}

func (c *metricProviderCatalog) ListByName() []string {
	repoList := make([]string, 0)

	for repoName, _ := range c.catalog {
		repoList = append(repoList, repoName)
	}
	return repoList
}
*/
