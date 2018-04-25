package sendsmsadapter

import (
	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	sendsms "github.com/StephaneBunel/sendsms/pkg/sms"

	"github.com/spf13/viper"
)

type (
	// Handle object
	Handle struct {
		config              *viper.Viper
		eventChan           domain.AmEventChan
		smsService          sendsms.ISmsService
		smsTemplate         string
		recipientRepository domain.RecipientRepositoryer
		metric              domain.Metricer
	}
)
