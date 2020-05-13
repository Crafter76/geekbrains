package main

import "net/http"

func (serv *Server) HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	serv.lg.Infof("%s %s", r.Method, r.URL)
	w.Write([]byte("Hello world"))
}
