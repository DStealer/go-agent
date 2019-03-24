package main

import (
	"flag"
	"log"
	"org/dstealer/agent/client"
	"org/dstealer/agent/server"
)

var serverMode = flag.Bool("server", false, "run as server mode")

var clientMode = flag.Bool("client", false, "run as client mode")

var configPath = flag.String("config", "", "config path,eg. config.json")

func main() {
	flag.Parse()
	if *serverMode {
		log.Printf("run as server mode")
		server.Run(*configPath)
	} else if *clientMode {
		log.Printf("run as client mode")
		client.Run(*configPath)
	} else {
		flag.Usage()
	}
}
