package sendsmsadapter

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	sendsms "github.com/StephaneBunel/sendsms/pkg/sms"

	"github.com/spf13/viper"
)

type (
	SendSmsInteractor struct {
		config                     *viper.Viper
		eventChan                  domain.AmEventChan
		smsService                 sendsms.ISmsService
		smsTemplate                string
		recipientRepositoryService domain.IRecipientRepositoryer
	}
)
