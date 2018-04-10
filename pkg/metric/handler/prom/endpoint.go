package prom

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/romana/rlog"
)

func (pm *promMetric) Serve() {
	endpoint := pm.config.GetString("endpoint")
	if endpoint == "" {
		rlog.Warn("stat endpoint is nil. Using /metrics")
		endpoint = "/metrics"
	}
	rlog.Info("Adding prometheus metric endpoint:", endpoint)
	http.Handle(endpoint, prometheus.Handler())
	address := pm.config.GetString("address")
	if address == "" {
		rlog.Warn("metric (prometheus) address is nil. Using 127.0.0.1:1471")
		address = "127.0.0.1:1471"
	}
	rlog.Info("metric (prometheus) listening at:", address)
	rlog.Critical(http.ListenAndServe(address, nil))
	os.Exit(1)
}
