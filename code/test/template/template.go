package template

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"text/template"
)

type Config struct {
	ServicePort    string `json:"servicePort"`
	DbHost         string `json:"dbHost"`
	DbPort         int    `json:"dbPort"`
	DbUsername     string `json:"dbUsername"`
	DbPassword     string `json:"dbPassword"`
	RedisHost      string `json:"redisHost"`
	RedisPort      string `json:"redisPort"`
	RedisUsername  string `json:"redisUsername"`
	RedisPassword  string `json:"redisPassword"`
	VrmServiceUrl  string `json:"vrmServiceUrl"`
	VrcmServiceUrl string `json:"vrcmServiceUrl"`
}

var (
	err    error
	f      *os.File
	tmpl   *template.Template
	config Config

	help        string
	filePath    string
	fileName    string
	outFilePath string
	outFileName string
	project     string
)

// ---- main ----
func main() {
	flag.StringVar(&help, "help", "", "eg: ./WillCG -filePath /usr/local/vs_conf/ -fileName vrpm.json -project vrpm -outFilePath /usr/local/vrpm/envconfig/ -outFileName dev_config.yaml")
	flag.StringVar(&filePath, "filePath", "/usr/local/vs_conf/", "(*OPT) the origin config path, like '/usr/local/vs_conf/'")
	flag.StringVar(&fileName, "fileName", "", "(*MUST) the origin config path, like 'vpm.json'")
	flag.StringVar(&outFilePath, "outFilePath", "./envconfig/", "(*OPT) the out put file path you want, like './envconfig'")
	flag.StringVar(&outFileName, "outFileName", "dev_config.yaml", "(*OPT) the out put file name you want, like 'dev_config.yaml' ")
	flag.StringVar(&project, "project", "vrpm", "(*MUST) the project you own, like 'vrpm' ")
	flag.Parse()

	vp := viper.New()
	vp.AddConfigPath(filePath)
	if len(filePath) <= 0 {
		log.Println("Please give the filePath, like '-filePath /usr/local/'")
		os.Exit(0)
	}
	if len(fileName) <= 0 {
		log.Println("Please give the fileName, like '-fileName vpm.json'")
		os.Exit(0)
	}
	vp.SetConfigName(fileName)
	vp.SetConfigType("json")
	err = vp.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("ReadInConfig Fail, Caused:   %v\n", err))
		os.Exit(0)
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println(fmt.Errorf("Parse Fail, Caused:  %v\n", err))
		os.Exit(0)
	}
	if len(project) > 0 {
		if val, ok := ProjectList[project]; ok {
			tmpl, err = template.New("test").Parse(val)
		} else {
			log.Println("Give your project name like: -project vrpm ")
			os.Exit(0)
		}
	} else {
		tmpl, err = template.New("test").Parse(vrpm)
	}

	if err != nil {
		log.Fatal(fmt.Errorf("Template Generate Fail, Caused:   %v\n", err))
		os.Exit(0)
	}
	CreateDir(outFilePath)
	f, err = os.Create(outFilePath + outFileName)
	defer func() {
		f.Close()
	}()
	err = tmpl.Execute(f, config)
	if err != nil {
		log.Fatal(fmt.Errorf("Template Generate Fail:  %v\n", err))
		os.Exit(0)
	}
	log.Println("Generate file success")
}

// ---- judge dir ----

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常： %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常： %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}

// ---- template ----
var ProjectList = map[string]string{
	"vrpm": vrpm,
	"vrcm": vrcm,
}

const (
	vrpm = `dev:
  mysql:
    vrpm:
      Host: {{.DbHost}}
      Port: {{.DbPort}}
      User: {{.DbUsername}}
      Password: {{.DbPassword}}
      DataBase: vrpm
      ParseTime: True
      MaxIdleConns: 10
      MaxOpenConns: 30
      ConnMaxLifetime: 28800
      ConnMaxIdletime: 7200
  redis:
    will:
      host: {{.RedisHost}}:{{.RedisPort}}
      password: {{.RedisPassword}}
      database: 0
      maxIdleNum: 50
      maxActive: 5000
      maxIdleTimeout: 600
      connectTimeout: 1
      readTimeout: 2
  rocketmq:
    GroupName: test-rocket
    Topic: test-rocket
    Host:
      - 127.0.0.1:9876
    Retry: 3`

	vrcm = `settings:
  application:
    mode: dev
    host: 0.0.0.0
    name: vrcm
    port: 18000
    readtimeout: 1
    writertimeout: 2
    enabledp: false
  logger:
    path: temp/logs
    level: trace
    enableddb: false
  #允许访问的域（用于设置前端跨域）
  alloworigin: http://0.0.0.0:18000
  database:
    driver: mysql
    source: {{.DbUsername}}:{{.DbPassword}}@tcp({{.DbHost}}:{{.DbPort}})/vrcm?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
  extend:
    serviceconfig:
      vrmurl: {{.VrmServiceUrl}}`
)
