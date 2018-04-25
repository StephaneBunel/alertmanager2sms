package am2sms

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/romana/rlog"
)

// CreateEventChan returns a new AmEvent channel
func CreateEventChan(cnf *appconfig.AppConfig, metric domain.Metricer) domain.AmEventChan {
	qlen := cnf.Viper.GetInt(cnf.AppName() + ".eventBufferSize")
	if qlen < 1 {
		qlen = 128
	}
	rlog.Debug("create eventChan with qlen = ", qlen)
	return domain.NewEventChan(qlen)
}
