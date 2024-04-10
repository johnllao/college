package main

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/johnllao/college/pkg/log"
	"github.com/johnllao/college/pkg/web"
)

const (
	BlankString = ""
)

func main() {
	var port, logpath, rootpath string
	flag.StringVar(&port, "port", "8080", "web port number")
	flag.StringVar(&rootpath, "root", BlankString, "root path")
	flag.StringVar(&logpath, "log", BlankString, "path of the logfile")
	flag.Parse()

	var err error

	log.Init(os.Stderr)
	log.SetLevel(log.LevelDebug)

	var l net.Listener
	if l, err = net.Listen("tcp", "localhost:"+port); err != nil {
		log.Fatal(err.Error())
	}
	defer l.Close()

	var server http.Server
	server.Handler = web.NewRootHandler(rootpath)

	log.Info("web server started. port: %s", port)
	if err = server.Serve(l); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Bye")
}
