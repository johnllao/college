package main

import (
	"flag"
	"html/template"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/johnllao/college/pkg/log"
	"github.com/johnllao/college/pkg/web"
)

const (
	BlankString = ""
)

type indexHandler struct {
	rootpath string
}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var t *template.Template
	t, err = template.ParseFiles(
		filepath.Join(h.rootpath, "index.html"),
		filepath.Join(h.rootpath, "head.html"),
		filepath.Join(h.rootpath, "header.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	if err = t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
}

type companiesHandler struct {
	rootpath string
}

func (h *companiesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var t *template.Template
	t, err = template.ParseFiles(
		filepath.Join(h.rootpath, "companies.html"),
		filepath.Join(h.rootpath, "head.html"),
		filepath.Join(h.rootpath, "header.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	if err = t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
}

func main() {
	var datapath, port, logpath, rootpath string
	flag.StringVar(&port, "port", "8080", "web port number")
	flag.StringVar(&rootpath, "root", BlankString, "root path")
	flag.StringVar(&logpath, "log", BlankString, "path of the logfile")
	flag.StringVar(&datapath, "data", BlankString, "path of the datafile")
	flag.Parse()

	var err error

	log.Init(os.Stderr)
	log.SetLevel(log.LevelDebug)

	var l net.Listener
	if l, err = net.Listen("tcp", "localhost:"+port); err != nil {
		log.Fatal(err.Error())
	}
	defer l.Close()

	var h = web.NewRootHandler(rootpath)
	h.IndexHandler = &indexHandler{rootpath: rootpath}
	h.Handle("/companies.html", &companiesHandler{rootpath: rootpath})

	var server http.Server
	server.Handler = h

	log.Info("web server started. port: %s", port)
	if err = server.Serve(l); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Bye")
}
