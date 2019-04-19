package composite

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	fmt.Println(WorkDir)
}

func TestReadPid(t *testing.T) {
	pid, err := ReadPid("./pid")
	if err != nil {
		t.Fatal("failed", err)
	} else {
		t.Log("pid:", pid)
	}
}

func TestWritePid(t *testing.T) {
	err := WritePid("./pid")
	if err != nil {
		t.Fatal(err)
	}
}

func TestHostname(t *testing.T) {
	hostname := Hostname("abc")
	t.Log("hostname:", hostname)
}
func TestMd5(t *testing.T) {
	md5str, err := Md5("/home/DStealer/Workspace/go-projects/go-agent/src/org/dstealer/agent/README.md")
	if err != nil {
		t.Fatal("failed", err)
	} else {
		t.Log("md5:", md5str)
	}
}
