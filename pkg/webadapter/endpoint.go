package webadapter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/StephaneBunel/alertmanager2sms/pkg/domain"

	"github.com/romana/rlog"
)

func (srv *WebserviceHandler) getHome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK."))
}

func (srv *WebserviceHandler) getNewAmEvent(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte((`{ "status": "ERROR", "message:", "You must use POST verb with this endpoint." }`)))
}

func (srv *WebserviceHandler) postNewAmEvent(w http.ResponseWriter, req *http.Request) {
	// params := mux.Vars(req)
	rlog.Debug("Connection from:", req.RemoteAddr)

	var event domain.AmEvent
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		rlog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf(`{ "status": "ERROR", "message:", "%s" }`, err.Error())
		w.Write([]byte(msg))
		return
	}
	rlog.Trace(1, "Event:", event)

	// Push event in channel (non blocking)
	select {
	case srv.eventChan <- &event:
	default:
		rlog.Error("Cannot push new event in channel. Try to raise eventBufferSize config parameter and restart.")
	}
}
