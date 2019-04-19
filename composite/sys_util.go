package composite

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var WorkDir string

func init() {
	path, err := pwd()
	if err == nil {
		WorkDir = path
	} else {
		WorkDir = "./"
	}
}

//当前运行目录
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

//从文件读取Pid
func ReadPid(filepath string) (pid int, err error) {
	_, err = os.Stat(filepath)
	if err != nil {
		return -1, err
	}
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(bytes))
}

//写入pid
func WritePid(filepath string) (err error) {
	_, err = os.Stat(filepath)
	if err != nil && os.IsExist(err) {
		return errors.New("pid exist")
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d", os.Getpid()))
	return err
}

//获取主机名
func Hostname(defaultHostname string) string {

	if strings.TrimSpace(defaultHostname) != "" {
		return strings.TrimSpace(defaultHostname)
	}

	name, err := os.Hostname()
	if err == nil {
		return name
	}

	addrs, err := net.InterfaceAddrs()

	if err == nil {
		return "unknown"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "unknown"
}

func Md5(filePath string) (md5str string, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	md5Inst := md5.New()
	for {
		buf := make([]byte, 8192)
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else {
			_, _ = io.WriteString(md5Inst, string(buf[0:n]))
		}
	}
	return base64.StdEncoding.EncodeToString(md5Inst.Sum(nil)), nil
}

func Md5SumCheck(workdir, md5file string) bool {
	cmd := exec.Command("md5sum", "-c", md5file)
	cmd.Dir = workdir
	err := cmd.Run()
	if err != nil {
		log.Printf("oops,cd %s; md5sum -c %s failed", workdir, md5file)
		return false
	}
	return true
}
