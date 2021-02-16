package webserver

import (
	"net/http"
)

func (s *WebServer) TimeGetPeriods(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TimeInsertPeriod(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TimeDeletePeriod(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

