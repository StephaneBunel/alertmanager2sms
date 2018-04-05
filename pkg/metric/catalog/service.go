package catalog

import (
	"strings"
	"sync"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
)

var (
	catalogMu sync.RWMutex
	catalog   = make(metricCatalog)
)

// Default returns the default metrics handlers catalog
func Default() MetricHandlerCataloger {
	return catalog
}

func (c metricCatalog) Register(factory func() domain.Metricer, info domain.MetricHandlerInfo) {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	var handlerName = strings.ToLower(info.Name)
	if handlerName == "" {
		panic("metric catalog: handler name is empty")
	}
	if factory == nil {
		panic("metric catalog: Factory func() is nil")
	}
	if _, duplicate := c[handlerName]; duplicate == true {
		panic("recipient repository catalog: Register() called twice for recipient repository: " + handlerName)
	}
	c[handlerName] = factory
}

func (c metricCatalog) Exists(name string) bool {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	_, exists := c[strings.ToLower(name)]
	return exists
}

func (c metricCatalog) New(name string) domain.Metricer {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	rr, exists := c[strings.ToLower(name)]
	if exists {
		return rr()
	}
	return nil
}

func (c metricCatalog) ListByName() []string {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	repoList := make([]string, 0)
	for repoName := range c {
		repoList = append(repoList, repoName)
	}
	return repoList
}
