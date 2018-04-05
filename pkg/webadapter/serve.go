package webadapter

import (
	"net/http"

	"github.com/romana/rlog"
)

func (srv *WebserviceHandle) Serve() {
	endpoint := "/"
	rlog.Info("Adding home endpoint: GET", endpoint)
	srv.router.HandleFunc(endpoint, srv.getHome).Methods("GET")

	endpoint = srv.config.GetString("endpoint")
	rlog.Info("Adding alertmanager endpoint: POST", endpoint)
	srv.router.HandleFunc(endpoint, srv.postNewAmEvent).Methods("POST")

	rlog.Info("Adding alertmanager endpoint: GET", endpoint)
	srv.router.HandleFunc(endpoint, srv.getNewAmEvent).Methods("GET")

	address := srv.config.GetString("address")
	rlog.Info("REST API listening at", address)
	rlog.Critical(http.ListenAndServe(address, srv.router))
}
