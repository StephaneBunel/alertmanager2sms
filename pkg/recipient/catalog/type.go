package catalog

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
)

type (
	RecipientRepositoryCataloger interface {
		Register(func() domain.RecipientRepositoryer, domain.RecipientRepositoryHandlerInfo)
		Exists(name string) bool
		New(name string) domain.RecipientRepositoryer
		ListByName() []string
	}

	repositoryCatalog map[string]func() domain.RecipientRepositoryer
)
