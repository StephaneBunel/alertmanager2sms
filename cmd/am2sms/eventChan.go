package main

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/appconfig"
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/romana/rlog"
)

func CreateEventChan(cnf *appconfig.AppConfig) domain.AmEventChan {
	qlen := cnf.Viper.GetInt(cnf.AppName() + ".eventBufferSize")
	if qlen < 1 {
		qlen = 128
	}
	rlog.Debug("create eventChan with qlen = ", qlen)
	return domain.NewEventChan(qlen)
}
