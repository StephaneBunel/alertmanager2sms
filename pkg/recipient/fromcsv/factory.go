package fromcsv

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/recipient"
)

func New() domain.IRecipientRepositoryer {
	return new(csvRecipientRepository)
}

func init() {
	recipient.RepositoryCatalog().Register(New, (&csvRecipientRepository{}).Info())
}
