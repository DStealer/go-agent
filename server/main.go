package server

import "log"

func Run(configPath string) {
	err := Parse(configPath)
	if err != nil {
		log.Fatal("parse config file failed", err)
	}
}
