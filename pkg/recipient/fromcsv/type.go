package fromcsv

import (
	"container/list"

	"github.com/spf13/viper"
)

type (
	csvRecipientRepository struct {
		conf            *viper.Viper
		filename        string
		linkedRecipient *list.List
	}
)
