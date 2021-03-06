package fromcsv

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/recipient/catalog"
)

// New returns a CSV implementation of a recipient repository
func New() domain.RecipientRepositoryer {
	return new(csvRepositoryHandle)
}

func init() {
	catalog.Default().Register(New, (&csvRepositoryHandle{}).Info())
}
