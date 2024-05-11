package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	var s = make(chan os.Signal, 1)
	signal.Notify(s)
	log.Println(<-s)
}
