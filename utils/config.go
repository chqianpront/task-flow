package utils

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type ServerConf struct {
	Server struct {
		Name string
		Port int
		Ip   string
	}
}

const (
	configDir       = "conf"
	confFilePrefix  = "application-"
	confFilePostfix = ".yml"
)

var appDirPath string

func init() {
	log.Printf("set default config")
	appDirPath, _ = os.Getwd()
}
func GetConf(envType string) *ServerConf {
	configFilePath := path.Join(appDirPath, configDir, confFilePrefix+envType+confFilePostfix)
	log.Printf("config file path: %s\n", configFilePath)
	fbytes, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Printf("read config file error: %s\n", err.Error())
		return nil
	}
	serverConf := &ServerConf{}
	yaml.Unmarshal(fbytes, serverConf)
	return serverConf
}
