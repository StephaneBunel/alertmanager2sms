package catalog

import (
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
