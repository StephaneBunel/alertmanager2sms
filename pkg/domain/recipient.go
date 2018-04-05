package domain

import "github.com/spf13/viper"

type (
	Recipient struct {
		Name         string
		PhoneNumbers []string
		Comment      string
		Tags         []string
	}

	RecipientList []*Recipient

	RecipientRepositoryHandlerInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}

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
