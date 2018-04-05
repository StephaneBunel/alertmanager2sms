package catalog

import (
	"strings"
	"sync"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/romana/rlog"
)

var (
	catalogMu sync.RWMutex
	catalog   = make(repositoryCatalog)
)

// Default returns the default recipients repositories catalog
func Default() RecipientRepositoryCataloger {
	return catalog
}

func (c repositoryCatalog) Register(factory func() domain.RecipientRepositoryer, info domain.RecipientRepositoryHandlerInfo) {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	var engineName = strings.ToLower(info.Name)
	if engineName == "" {
		panic("recipient repository catalog: name is empty")
	}
	if factory == nil {
		panic("recipient repository catalog: Factory func() is nil")
	}
	if _, duplicate := c[engineName]; duplicate == true {
		panic("recipient repository catalog: Register() called twice for recipient repository: " + engineName)
	}
	c[engineName] = factory
}

func (c repositoryCatalog) Exists(name string) bool {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	_, exists := c[strings.ToLower(name)]
	return exists
}

func (c repositoryCatalog) New(name string) domain.RecipientRepositoryer {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	rr, exists := c[strings.ToLower(name)]
	if exists {
		return rr()
	}
	rlog.Error("recipient repository [", name, "] not found in catalog")
	return nil
}

func (c repositoryCatalog) ListByName() []string {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	repoList := make([]string, 0)
	for repoName := range c {
		repoList = append(repoList, repoName)
	}
	return repoList
}
