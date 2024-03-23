package main

import (
	"log"
	"net/http"
)

type RootHandler struct {
	rootpath string

	fs http.Handler
}

func NewRootHandler(path string) *RootHandler {
	var h RootHandler
	h.rootpath = path

	log.Printf("initializing file server. path: %s", h.rootpath)
	h.fs = http.FileServer(http.Dir(h.rootpath))

	return &h
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	h.fs.ServeHTTP(w, r)
}
