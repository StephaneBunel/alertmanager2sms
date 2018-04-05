package sendsmsadapter

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	sendsms "github.com/StephaneBunel/sendsms/pkg/sms"

	"github.com/spf13/viper"
)

type (
	SendsmsAdapterHandle struct {
		config              *viper.Viper
		eventChan           domain.AmEventChan
		smsService          sendsms.ISmsService
		smsTemplate         string
		recipientRepository domain.RecipientRepositoryer
	}
)
