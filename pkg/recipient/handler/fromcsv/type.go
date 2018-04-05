package fromcsv

import (
	"container/list"

	"github.com/spf13/viper"
)

type (
	csvRepositoryHandle struct {
		conf            *viper.Viper
		filename        string
		linkedRecipient *list.List
	}
)
