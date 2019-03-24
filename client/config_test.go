package client

import (
	"fmt"
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	err := Parse("resources/config.json")
	if err != nil {
		log.Fatal("read file failed. ", err)
	}
	globalConfig := Config()
	fmt.Println(globalConfig)
}
