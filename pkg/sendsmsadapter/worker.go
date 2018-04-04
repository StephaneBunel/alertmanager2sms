package sendsmsadapter

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/romana/rlog"
)

func (s *SendSmsInteractor) Worker() error {
	for {
		select {
		case event := <-s.eventChan:
			rlog.Debug(event)
			recipient := event.Receiver
			rlog.Debug("recipient =", recipient)
			recipients := s.recipientRepositoryService.FindByName(recipient)
			if len(recipients) == 0 {
				rlog.Debug("recipient(s) not found")
				break
			}
			phones := recipientsToPhones(recipients)
			rlog.Debug("Recipient phones:", phones)
			text := TemplateAmEvent(event, s.smsTemplate)
			err := s.smsService.SendRaw(text, phones...)
			if err != nil {
				rlog.Error(err)
			}
		}
	}
}

func recipientsToPhones(recipients domain.RecipientList) []string {
	phones := make([]string, 0)
	for _, recipient := range recipients {
		phones = append(phones, recipient.PhoneNumbers...)
	}
	return phones
}

func AmEventToString(event *domain.AmEvent) string {
	txt := ""
	txt += fmt.Sprintf("[%s] %d alerts:\n",
		strings.ToUpper(event.Status),
		len(event.Alerts))
	for i, alert := range event.Alerts {
		txt += AmAlertToString(i, len(event.Alerts), alert)
	}

	return txt
}

func AmAlertToString(index int, total int, alert domain.AmAlert) string {
	return fmt.Sprintf("(%d/%d): %s - %s\n%s\nBegin: %s",
		index+1, total,
		strings.ToUpper(alert.Status),
		strings.ToUpper(alert.Labels["severity"]),
		alert.Annotations["description"],
		alert.StartsAt)
}

func TemplateAmEvent(event *domain.AmEvent, tpl string) string {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
		"add": func(a, b int) int {
			return a + b
		},
	}

	t, err := template.New("body").Funcs(funcMap).Parse(tpl)
	if err != nil {
		rlog.Error(err)
		return ""
	}
	buf := bytes.NewBufferString("")
	t.Execute(buf, event)
	return buf.String()
}
