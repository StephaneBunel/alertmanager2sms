package webadapter

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func New(config *viper.Viper, eventChan domain.AmEventChan) *WebserviceHandler {
	w := new(WebserviceHandler)
	w.config = config
	w.eventChan = eventChan
	w.router = mux.NewRouter()
	return w
}
