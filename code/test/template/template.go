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
	flag.StringVar(&templates, ".", "", "(*MUST) the template you want")
	flag.StringVar(&outFilePath, "outFilePath", "./envconfig", "(*OPT) the out put file path you want like './envconfig'")
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
