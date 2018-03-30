package domain

import "github.com/spf13/viper"

type (
	IRecipientRepository interface {
		Config(*viper.Viper) error
		Add(*Recipient) error
		FindByName(name string) RecipientList
		Info() RecipientRepositoryInfo
	}

	RecipientRepositoryInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}
)
