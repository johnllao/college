package main

import (
	"flag"
	"log"
	"net"
	"net/http"
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

	var l net.Listener
	if l, err = net.Listen("tcp", "localhost:"+port); err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	var server http.Server
	server.Handler = NewRootHandler(rootpath)

	log.Printf("web server started. port: %s", port)
	if err = server.Serve(l); err != nil {
		log.Fatal(err)
	}

	log.Print("Bye")
}
