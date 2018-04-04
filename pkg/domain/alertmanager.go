package domain

import "github.com/romana/rlog"

type (
	// AlertManagerEvent represents the JSON struct that is POST'd to a web_hook
	// received from Prometheus' Alertmanager.  There are other fields in the
	// JSON blob which are not included here.
	AmEvent struct {
		Receiver          string            `json:"receiver"` // Group to send messages to
		Status            string            `json:"status"`
		Alerts            []AmAlert         `json:"alerts"`
		GroupLabels       map[string]string `json:"groupLabels"`
		CommonLabels      map[string]string `json:"commonLabels"`
		CommonAnnotations map[string]string `json:"commonAnnotations"`
		ExternalURL       string            `json:"externalURL"`
		Version           string            `json:"version"`
		GroupKey          int               `json:"groupKey"`
	}

	// Alert represents an individual alert from Prometheus and included in the
	// JSON blob POST'd via the Alertmanager.
	AmAlert struct {
		Status       string            `json:"status"`
		Labels       map[string]string `json:"labels"`
		Annotations  map[string]string `json:"annotations"`
		StartsAt     string            `json:"startsAt"`
		EndsAt       string            `json:"endsAt"`
		GeneratorURL string            `json:"generatorURL"`
		Json         string            `json:"-"` // Json is not from the alert JSON but holds a JSON formatted string
		// of this alert.  It is not the same JSON as originally passed in.
	}

	AmEventChan chan *AmEvent
)

// NewEvantChan returns a new Alertmanger Event channel
func NewEventChan(qlen int) AmEventChan {
	if qlen < 1 {
		qlen = 128
		rlog.Info("set alertmanger event queue channel to", qlen, "because it is lower than 1.")
	}
	return make(AmEventChan, qlen)
}
