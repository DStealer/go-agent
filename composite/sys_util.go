package composite

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var WorkDir string

func init() {
	path, err := pwd()
	if err == nil {
		WorkDir = path
	} else {
		WorkDir = "."
	}
}

func pwd() (path string, err error) {
	path, err = exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return "", err
	}
	lastPathSep := strings.LastIndexByte(path, os.PathSeparator)
	if lastPathSep > 0 {
		return string(path[0 : lastPathSep+1]), nil
	} else {
		return path, nil
	}
}
