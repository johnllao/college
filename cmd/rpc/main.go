package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"strconv"

	"github.com/johnllao/college/cmd/rpc/args"
	"github.com/johnllao/college/cmd/rpc/ops"
)

func main() {

	var clientmode, servermode bool
	var cmd string
	var port int
	flag.BoolVar(&clientmode, "client", false, "client mode")
	flag.StringVar(&cmd, "cmd", "", "name of the command")
	flag.IntVar(&port, "port", 9090, "service port")
	flag.BoolVar(&servermode, "server", false, "server mode")
	flag.Parse()

	if clientmode && servermode {
		log.Fatal("cannot enable both client and server mode")
	}

	if clientmode {
		startclient(port, cmd)
	}

	if servermode {
		startserver(port)
	}
}

func startclient(port int, cmd string) {
	var err error
	var cli *rpc.Client
	if cli, err = rpc.Dial("tcp", "localhost:"+strconv.Itoa(port)); err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	switch cmd {
	case "ping":
		var a int
		var reply args.PingReply
		if err = cli.Call("ServerOps.Ping", &a, &reply); err != nil {
			log.Fatal(err)
		}
		log.Print(reply)
	case "status":
		var a int
		var reply args.StatusReply
		if err = cli.Call("ServerOps.Status", &a, &reply); err != nil {
			log.Fatal(err)
		}
		log.Print(reply)
	}
}

func startserver(port int) {
	var err error

	var s = rpc.NewServer()
	s.RegisterName("ServerOps", new(ops.ServerOps))

	var l net.Listener
	if l, err = net.Listen("tcp", "localhost:"+strconv.Itoa(port)); err != nil {
		log.Print(err)
	}
	defer l.Close()

	log.Print("service started.")
	s.Accept(l)
}
