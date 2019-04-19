package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var config *GlobalConfig
var lock = new(sync.RWMutex)

type AgentConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type GlobalConfig struct {
	Pid        string         `json:"pid"`
	Http       *HttpConfig    `json:"http"`
	TarballDir string         `json:"tarballDir"`
	Agents     []*AgentConfig `json:"agents"`
}

/*
获取全局配置文集
*/
func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

/*
解析配置文件
*/
func Parse(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	log.Print("read config from:", file.Name())
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	err = json.Unmarshal(bytes, &config)
	return err
}
