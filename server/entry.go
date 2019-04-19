package server

import (
	"log"
	"net/http"
)

func Run(configPath string) {
	err := Parse(configPath)
	if err != nil {
		log.Fatal("parse config file failed", err)
	}
	config = Config()
	log.Println("load the config:", config)
	if !config.Http.Enabled {
		log.Println("http is disable")
		return
	}
	if config.Http.Listen == "" {
		log.Fatalln("listen is blank")
		return
	}

	healthFunc()
	heartbeatFunc()
	configFunc()
	status()

	server := http.Server{Addr: config.Http.Listen, MaxHeaderBytes: 1 << 30}
	log.Println("listening:", config.Http.Listen)
	log.Fatalln("http server running failed", server.ListenAndServe())
}
