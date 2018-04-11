package sendsmsadapter

import (
	"bytes"
	"strings"
	"text/template"
	"time"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/romana/rlog"
)

func (s *SendsmsAdapterHandle) Worker() error {
	for {
		select {
		case event := <-s.eventChan:
			t0 := time.Now()
			rlog.Debug(event)
			recipient := event.Receiver
			rlog.Debug("recipient =", recipient)
			recipients := s.recipientRepository.FindByName(recipient)
			if len(recipients) == 0 {
				rlog.Debug("recipient(s) not found")
			} else {
				phones := recipientsToPhones(recipients)
				rlog.Debug("Recipient phones:", phones)
				text := TemplateAmEvent(event, s.smsTemplate)
				err := s.smsService.SendRaw(text, phones...)
				if err != nil {
					rlog.Error(err)
				}
				s.metric.IncCounter("sms_sent_total")
			}
			tt := time.Since(t0)
			rlog.Debug("sendsms duration:", tt.String())
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
