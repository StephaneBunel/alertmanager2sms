package fromcsv

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/recipient/catalog"
)

func New() domain.IRecipientRepositoryer {
	return new(csvRecipientRepository)
}

func init() {
	catalog.Default().Register(New, (&csvRecipientRepository{}).Info())
}
