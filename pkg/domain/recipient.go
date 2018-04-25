package domain

import "github.com/spf13/viper"

type (
	// Recipient represent the minimal informations about a recipient
	Recipient struct {
		Name         string
		PhoneNumbers []string
		Comment      string
		Tags         []string
	}

	// RecipientList is a list of recipients
	RecipientList []*Recipient

	// RecipientRepositoryHandlerInfo must be passed by each recipient repository when it registers itself.
	RecipientRepositoryHandlerInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}

	// RecipientRepositoryer is a recipient repository interface
	RecipientRepositoryer interface {
		Config(*viper.Viper) error
		Add(*Recipient) error
		FindByName(name string) RecipientList
		Info() RecipientRepositoryHandlerInfo
	}
)

// NewRecipient returns a new recipient object
func NewRecipient() *Recipient {
	return new(Recipient)
}
