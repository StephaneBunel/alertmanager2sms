package recipient

import (
	"strings"
	"sync"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
)

type (
	IRecipientRepositoryCatalog interface {
		Register(func() domain.IRecipientRepositoryer, domain.RecipientRepositoryInfo)
		Exists(name string) bool
		New(name string) domain.IRecipientRepositoryer
		ListByName() []string
	}

	repositoryCatalog map[string]func() domain.IRecipientRepositoryer
)

var (
	catalogMu sync.RWMutex
	catalog   repositoryCatalog = make(repositoryCatalog)
)

func RepositoryCatalog() IRecipientRepositoryCatalog {
	return catalog
}

func (c repositoryCatalog) Register(factory func() domain.IRecipientRepositoryer, info domain.RecipientRepositoryInfo) {
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

func (c repositoryCatalog) New(name string) domain.IRecipientRepositoryer {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	rr, exists := c[strings.ToLower(name)]
	if exists {
		return rr()
	}
	return nil
}

func (c repositoryCatalog) ListByName() []string {
	catalogMu.Lock()
	defer catalogMu.Unlock()
	repoList := make([]string, 0)
	for repoName, _ := range c {
		repoList = append(repoList, repoName)
	}
	return repoList
}
