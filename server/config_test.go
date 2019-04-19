package server

import (
	"fmt"
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	err := Parse("/home/DStealer/Workspace/go-projects/go-agent/src/org/dstealer/agent/resources/server.json")
	if err != nil {
		log.Fatal("read file failed. ", err)
	}
	globalConfig := Config()
	fmt.Println(globalConfig)
}
