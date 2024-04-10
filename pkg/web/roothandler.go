package web

import (
	"net/http"

	"github.com/johnllao/college/pkg/log"
)

type RootHandler struct {
	rootpath string

	fs  http.Handler
	mux *http.ServeMux

	IndexHandler http.Handler
}

func NewRootHandler(path string) *RootHandler {
	var h RootHandler
	h.rootpath = path

	log.Info("initializing file server. path: %s", h.rootpath)
	h.fs = http.FileServer(http.Dir(h.rootpath))
	h.mux = http.NewServeMux()
	return &h
}

func (h *RootHandler) Handle(pattern string, handler http.Handler) {
	h.mux.Handle(pattern, handler)
}

func (h *RootHandler) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.mux.HandleFunc(pattern, handler)
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		muxh      http.Handler
		pattern   string
		serveHTTP = h.fs.ServeHTTP
	)
	if muxh, pattern = h.mux.Handler(r); pattern != "" {
		serveHTTP = muxh.ServeHTTP
	}
	if h.IndexHandler != nil && (r.URL.Path == "/" || r.URL.Path == "/index.html") {
		serveHTTP = h.IndexHandler.ServeHTTP
	}
	serveHTTP(w, r)
}
