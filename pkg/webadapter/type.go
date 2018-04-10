package webadapter

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type (
	WebserviceHandle struct {
		config    *viper.Viper
		router    *mux.Router
		eventChan domain.AmEventChan
		metricer  domain.Metricer
	}
)
