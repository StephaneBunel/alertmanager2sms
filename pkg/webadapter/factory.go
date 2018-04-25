package webadapter

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// New returns a new webadapter
func New(config *viper.Viper, eventChan domain.AmEventChan, metric domain.Metricer) *WebserviceHandle {
	w := new(WebserviceHandle)
	w.config = config
	w.eventChan = eventChan
	w.router = mux.NewRouter()
	w.metricer = metric
	return w
}
