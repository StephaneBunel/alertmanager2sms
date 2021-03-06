package webadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"

	"github.com/romana/rlog"
)

func (srv *WebserviceHandle) getHome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK."))
}

func (srv *WebserviceHandle) getNewAmEvent(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte((`{ "status": "ERROR", "message:", "You must use POST verb with this endpoint." }`)))
}

func (srv *WebserviceHandle) postNewAmEvent(w http.ResponseWriter, req *http.Request) {
	t0 := time.Now()
	// params := mux.Vars(req)
	rlog.Debug("Connection from:", req.RemoteAddr)
	srv.metricer.IncCounter("web_request_total")

	var event domain.AmEvent
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		rlog.Error(err)
		srv.metricer.IncCounter("web_request_decode_error")
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf(`{ "status": "ERROR", "message:", "%s" }`, err.Error())
		w.Write([]byte(msg))
		tt := time.Since(t0)
		rlog.Debug("web request duration:", tt.String())
		return
	}
	rlog.Trace(1, "Event:", event)

	// Push event in channel (non blocking)
	select {
	case srv.eventChan <- &event:
	default:
		rlog.Error("Cannot push new event in channel. Try to raise eventBufferSize config parameter and restart.")
		srv.metricer.IncCounter("event_dropped_total")
	}
	tt := time.Since(t0)
	rlog.Debug("web request duration:", tt.String())
}
