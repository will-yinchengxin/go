package oss

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	AccessKeyId     string `yaml:"accessKeyId,omitempty"`
	AccessKeySecret string `yaml:"accessKeySecret,omitempty"`
	Host            string `yaml:"host,omitempty"`
	UploadDir       string `yaml:"uploadDir,omitempty"`
	ExpireTime      int64  `yaml:"expireTime,omitempty"`
}

type ApplicationConfig struct {
	Oss    Config `yaml:"oss,omitempty"`
	Locker *sync.RWMutex
}

func (conf *ApplicationConfig) Lock() {
	conf.Locker.Lock()
}

func (conf *ApplicationConfig) Unlock() {
	conf.Locker.Unlock()
}

var AppConfig *ApplicationConfig

func SetConfig() {
	AppConfig = &ApplicationConfig{Locker: &sync.RWMutex{}}
	path, err := os.Getwd() // D:\Project\test\
	if err != nil {
		panic(err)
	}
	path = path + "\\oss\\environment\\" + "app.yaml"
	err = setCustomerConfig(path, AppConfig)
	if err != nil { //判断是否获取成功
		panic(err)
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		setCustomerConfig(path, AppConfig)
	})
}

// readConfigOfNacos init nacos dail information
func setCustomerConfig(path string, appConfig *ApplicationConfig) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	//将数据写入系统配置变量
	err = yaml.Unmarshal(content, appConfig)
	if err != nil {
		return
	}
	return
}
