package sendsmsadapter_test

import (
	"testing"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"
	"github.com/StephaneBunel/alertmanager2sms/pkg/sendsmsadapter"
)

var event = domain.AmEvent{
	Receiver: "sysadmin",
	Status:   "firing",
	Alerts: []domain.AmAlert{
		domain.AmAlert{
			Status: "firing",
			Labels: map[string]string{
				"alertname": "node_up",
				"instance":  "cc3-admin-1",
				"job":       "node_exporter",
				"monitor":   "packet/nj",
				"severity":  "critical"},
			Annotations: map[string]string{
				"description": "cc3-admin-1 of job node_exporter has been down for more than 5 minutes.",
				"summary":     "Instancecc3-admin-1 is down"},
			StartsAt:     "2017-10-31T09:10:44.021Z",
			EndsAt:       "0001-01-01T00:00:00Z",
			GeneratorURL: "https://cc3-admin.inpsn.net/prometheus/graph?g0.expr=up+%3D%3D+0&g0.tab=0", Json: ""}}, GroupLabels: map[string]string{"alertname": "node_up"}, CommonLabels: map[string]string{"alertname": "node_up", "instance": "cc3-admin-1", "job": "node_exporter", "monitor": "packet/nj", "severity": "critical"}, CommonAnnotations: map[string]string{"description": "cc3-admin-1 of job node_exporter has been down for more than 5 minutes.", "summary": "Instance cc3-admin-1 is down"}, ExternalURL: "https://cc3-admin.inpsn.net/alertmanager",
	Version:  "3",
	GroupKey: 4964808624378674382}

var tpl = `
{{ $numAlerts := len .Alerts }}Hello {{ .Receiver }},
Alertmanger raised {{ $numAlerts }} {{ .Status | ToUpper }} alert(s):
{{ range $index, $alert := .Alerts }}
({{ add $index 1 }}/{{ $numAlerts }}) {{ $alert.Status | ToUpper }} {{ $alert.Labels.severity | ToUpper }} {{ $alert.Labels.instance }}/{{ $alert.Labels.job }}
Begin: {{ $alert.StartsAt }}
{{ $alert.Annotations.description }}
{{ end }}
Thank you.
`

var expectedResults = `
Hello sysadmin,
Alertmanger raised 1 FIRING alert(s):

(1/1) FIRING CRITICAL cc3-admin-1/node_exporter
Begin: 2017-10-31T09:10:44.021Z
cc3-admin-1 of job node_exporter has been down for more than 5 minutes.

Thank you.
`

func TestTemplate(t *testing.T) {

	out := sendsmsadapter.TemplateAmEvent(&event, tpl)
	if out != expectedResults {
		t.Error(out)
	}
}
