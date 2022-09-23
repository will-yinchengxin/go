package template

package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"text/template"
)

type Config struct {
	ServicePort   string `json:"servicePort"`
	DbHost        string `json:"dbHost"`
	DbPort        int    `json:"dbPort"`
	DbUsername    string `json:"dbUsername"`
	DbPassword    string `json:"dbPassword"`
	RedisHost     string `json:"redisHost"`
	RedisPort     string `json:"redisPort"`
	RedisUsername string `json:"redisUsername"`
	RedisPassword string `json:"redisPassword"`
}

var defaultTemplates = `dev:
 mysql:
   will:
     Host: {{.DbHost}}
     Port: {{.DbPort}}
     User: {{.DbUsername}}
     Password: {{.DbPassword}}
     DataBase: vcfn
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
   Retry: 3
`
var (
	err    error
	f      *os.File
	tmpl   *template.Template
	config Config

	help        string
	filePath    string
	fileName    string
	templates   string
	outFilePath string
	outFileName string
)

func main() {
	flag.StringVar(&help, "help", "", "eg: ./WillCG -filePath ./ -fileName vrpm.json -templates xx....xx ")
	flag.StringVar(&filePath, "filePath", "/usr/local/vs_conf/", "(*OPT) the origin config path, like '/usr/local/vs_conf/'")
	flag.StringVar(&fileName, "fileName", "", "(*MUST) the origin config path, like 'vpm.json'")
	flag.StringVar(&templates, "templates", "", "(*MUST) the template you want")
	flag.StringVar(&outFilePath, "outFilePath", "./envconfig/", "(*OPT) the out put file path you want like './envconfig'")
	flag.StringVar(&outFileName, "outFileName", "dev_config.yaml", "(*OPT) the out put file name you want")
	flag.Parse()

	vp := viper.New()
	vp.AddConfigPath(filePath)
	if len(fileName) <= 0 {
		log.Println("Please give the fileName, like '-fileName vpm.json'")
		os.Exit(0)
	}
	vp.SetConfigName(fileName)
	vp.SetConfigType("json")
	err = vp.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("ReadInConfig Fail, Caused:\n\n   %v\n", err))
		os.Exit(0)
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println(fmt.Errorf("Parse Fail, Caused:\n\n   %v\n", err))
		os.Exit(0)
	}
	if len(templates) <= 0 {
		tmpl, err = template.New("test").Parse(defaultTemplates)
	} else {
		tmpl, err = template.New("test").Parse(templates)
	}
	if err != nil {
		log.Fatal(fmt.Errorf("Template Generate Fail, Caused:\n\n   %v\n", err))
		os.Exit(0)
	}
	f, _ = os.Create(outFilePath + outFileName)
	defer func() {
		f.Close()
	}()
	err = tmpl.Execute(f, config)
	if err != nil {
		log.Fatal(fmt.Errorf("Template Generate Fail:\n\n   %v\n", err))
		os.Exit(0)
	}
	log.Println("Generate file success")
}

/*
./WillCG -filePath ./  -fileName vrpm.json  -outFileName dev.yaml  -outFilePath ./    -templates "settings:
  application:
    # dev开发环境 test测试环境 prod线上环境
    mode: dev
    # 服务器ip，默认使用 0.0.0.0
    host: {{.DbHost}}
    # 服务名称
    name: testApp
    # 端口号
    port: 8000 # 服务端口号
    readtimeout: 1
    writertimeout: 2
    # 数据权限功能开关
    enabledp: false
  logger:
    # 日志存放路径
    path: temp/logs
    # 日志输出，file：文件，default：命令行，其他：命令行
    stdout: '' #控制台日志，启用后，不输出到文件
    # 日志等级, trace, debug, info, warn, error, fatal
    level: trace
    # 数据库日志开关
    enableddb: false
  jwt:
    # token 密钥，生产环境时及的修改
    secret: vrcm
    # token 过期时间 单位：秒
    timeout: 3600
  database:
    # 数据库类型 mysql, sqlite3, postgres, sqlserver
    # sqlserver: sqlserver://用户名:密码@地址?database=数据库名
    driver: mysql
    # 数据库连接字符串 mysql 缺省信息 charset=utf8&parseTime=True&loc=Local&timeout=1000ms
    source: root:1726asLYH@tcp(127.0.0.1:3306)/gaa?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
  gen:
    # 代码生成读取的数据库名称
    dbname: dbname
    # 代码生成是使用前端代码存放位置，需要指定到src文件夹，相对路径
    frontpath: ../vrcm-ui/src
  extend: # 扩展项使用说明
    demo:
      name: data
  cache:
    memory: ''
  queue:
    memory:
      poolSize: 100
  locker:
    redis:
"
*/
