package fromcsv

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
)

func New() domain.IRecipientRepository {
	return new(csvRecipientRepository)
}

func init() {
	domain.RecipientRepositoryCatalog().Add(New, (&csvRecipientRepository{}).Info())
}
