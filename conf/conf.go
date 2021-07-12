package conf

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

var (
	Conf              config // Conf 全局配置信息
	defaultConfigFile = "conf/conf.toml"
)

type config struct {
	LogLevel string `toml:"log_level"`

	Server server   `toml:"server"`
	DB     database `toml:"database"`
}

type server struct {
	Port int `toml:"port"`
}

type database struct {
	Driver   string `toml:"driver"`
	FileName string `toml:"fileName"`

	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	UserName string `toml:"username"`
	PassWord string `toml:"password"`
	Name     string `toml:"name"`
}

// InitConfig 加载配置文件
func InitConfig(configFile string) {
	if configFile == "" {
		configFile = defaultConfigFile
	}
	Conf = config{}

	if _, err := os.Stat(configFile); err != nil {
		panic(errors.New("配置文件错误:" + err.Error()))
	} else {
		log.Infof("加载配置文件:" + configFile)
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(errors.New("配置文件加载错误:" + err.Error()))
		}
		_, err = toml.Decode(string(configBytes), &Conf)
		if err != nil {
			panic(errors.New("配置文件编码错误:" + err.Error()))
		}
	}
	log.Debugf("config data:%v", Conf)
}
