package mock

import (
	"log"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	command := exec.Command("pwd")
	bytes, err := command.Output()
	if err != nil {
		log.Fatalf("failed to output %v", err)
	}
	log.Println("output:", string(bytes))
}
