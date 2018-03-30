package domain

import (
	"strings"

	"github.com/romana/rlog"
)

type (
	IRecipientRepositoryCatalog interface {
		Add(func() IRecipientRepository, RecipientRepositoryInfo)
		Exists(name string) bool
		NewByName(name string) IRecipientRepository
		ListByName() []string
	}

	recipientRepositoryCatalog struct {
		catalog map[string]func() IRecipientRepository
	}
)

var (
	recipientRepoCatalog = &recipientRepositoryCatalog{
		catalog: make(map[string]func() IRecipientRepository),
	}
)

func RecipientRepositoryCatalog() IRecipientRepositoryCatalog {
	return recipientRepoCatalog
}

func (c *recipientRepositoryCatalog) Add(factory func() IRecipientRepository, info RecipientRepositoryInfo) {
	c.catalog[strings.ToLower(info.Name)] = factory
}

func (c *recipientRepositoryCatalog) Exists(name string) bool {
	rlog.Debug(c)
	_, exists := c.catalog[strings.ToLower(name)]
	return exists
}

func (c *recipientRepositoryCatalog) NewByName(name string) IRecipientRepository {
	if c.Exists(name) {
		return c.catalog[strings.ToLower(name)]()
	}
	return nil
}

func (c *recipientRepositoryCatalog) ListByName() []string {
	repoList := make([]string, 0)

	for repoName, _ := range c.catalog {
		repoList = append(repoList, repoName)
	}
	return repoList
}
