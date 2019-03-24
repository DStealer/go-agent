package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var config *GlobalConfig
var lock = new(sync.RWMutex)

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type GlobalConfig struct {
	Debug bool        `json:"debug"`
	Pid   string      `json:"pid"`
	Http  *HttpConfig `json:"http"`
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
