package sendsmsadapter

import (
	"os"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	_ "github.com/StephaneBunel/sendsms/pkg/provider"
	sendsms "github.com/StephaneBunel/sendsms/pkg/sms"

	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

const defaultSmsTemplate = `{{ $numAlerts := len .Alerts }}Hello {{ .Receiver }},
Alertmanger raised {{ $numAlerts }} {{ .Status | ToUpper }} alert(s):
{{ range $index, $alert := .Alerts }}
({{ add $index 1 }}/{{ $numAlerts }}) {{ $alert.Status | ToUpper }} {{ $alert.Labels.severity | ToUpper }} {{ $alert.Labels.instance }}/{{ $alert.Labels.job }}
Begin: {{ $alert.StartsAt }}
{{ $alert.Annotations.description }}
{{ end }}
Thank you.
`

func New(configSms *viper.Viper, configSendSms *viper.Viper, eventChan domain.AmEventChan, rrs domain.IRecipientRepositoryer) *SendSmsInteractor {
	if configSms == nil {
		rlog.Error("configSms is nil")
		os.Exit(1)
	}

	if configSendSms == nil {
		rlog.Error("configSendSms is nil")
		os.Exit(1)
	}

	ss := new(SendSmsInteractor)
	ss.config = configSms
	ss.eventChan = eventChan
	ss.recipientRepositoryService = rrs

	confKey := "profile"
	profileName := ss.config.GetString(confKey)
	if profileName == "" {
		rlog.Warn(confKey, "is not defined. Fallback to defaut profile")
		profileName = "default"
	}
	rlog.Debug("profileName =", profileName)

	provider := selectAndConfigureProviderFromProfile(configSendSms, profileName)
	rlog.Debug("provider =", provider)

	ss.smsService = sendsms.NewSmsService(provider)
	rlog.Debug("smsService =", ss.smsService)

	confKey = "template"
	ss.smsTemplate = ss.config.GetString(confKey)
	if ss.smsTemplate == "" {
		ss.smsTemplate = defaultSmsTemplate
	}

	return ss
}

func selectAndConfigureProviderFromProfile(config *viper.Viper, profileName string) sendsms.IProviderService {
	confKey := "profiles." + profileName + ".provider"
	profileProviderName := config.GetString(confKey)
	if profileProviderName == "" {
		rlog.Error(confKey, "Provider must be defined.")
		os.Exit(1)
	}
	rlog.Debug("Provider is", profileProviderName)
	provider := sendsms.GetProviderRepository().FindByName(profileProviderName)
	if provider == nil {
		rlog.Error("Provider: ", profileProviderName, "not found.")
		os.Exit(1)
	}
	confKey = "profiles." + profileName + ".providerConfig"
	providerConfig := config.Sub(confKey)
	if providerConfig == nil {
		rlog.Error("providerConfig is nil")
		os.Exit(1)
	}
	provider.Config(providerConfig)
	return provider
}
