package oss

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"sync"
)

type OssConfig struct {
	AccessKeyId     string `yaml:"accessKeyId,omitempty"`
	AccessKeySecret string `yaml:"accessKeySecret,omitempty"`
	Host            string `yaml:"host,omitempty"`
	UploadDir       string `yaml:"uploadDir,omitempty"`
	ExpireTime      int64  `yaml:"expireTime,omitempty"`
	BucketName      string `yaml:"bucketName,omitempty"`
	Endpoint        string `yaml:"endpoint,omitempty"`
	DwnLoadPartNum  int64  `yaml:"downLoadPartNum,omitempty"`
}

type ObsConfig struct {
	BucketName      string `yaml:"bucketName,omitempty"`
	ObjectKey       string `yaml:"objectKey,omitempty"`
	AK              string `yaml:"ak,omitempty"`
	SK              string `yaml:"sk,omitempty"`
	EndPoint        string `yaml:"endPoint,omitempty"`
	ExpireTime      int32  `yaml:"expireTime,omitempty"`
	Host            string `yaml:"host,omitempty"`
	Dir             string `json:"dir"`
	DownLoadPartNum int64  `yaml:"downLoadPartNum,omitempty"`
	DomainName      string `yaml:"domainName,omitempty"`
	UserName        string `yaml:"userName,omitempty"`
	UserPassword    string `yaml:"userPassword,omitempty"`
}

type ApplicationConfig struct {
	Oss    OssConfig `yaml:"oss,omitempty"`
	Obs    ObsConfig `yaml:"obs,omitempty"`
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

	// 动态更新配置文件信息
	v := viper.New()
	v.SetConfigFile(path)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		setCustomerConfig(path, AppConfig)
	})
}

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
