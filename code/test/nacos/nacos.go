package nacos

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"log"
	"sync"
)

/*
	中文文档：https://nacos.io/zh-cn/docs/what-is-nacos.html
*/

var EnvDev = &Conn{
	Addr:         "172.16.252.99",
	SlaveAddr:    "172.16.252.99",
	Namespace:    "70cd7d2e-85f2-4756-847c-a224bf89dc35",
	Path:         "/nacos",
	Port:         8848,
	CacheDir:     "storage/cache",
	LogsDir:      "storage/logs",
	TimeoutMs:    2 * 1000,
	BeatInterval: 10 * 1000,
}

const (
	APP_GROUP = "APPID_100001"
	DataId    = "system.config"
)

func FetchCoreConfig() bool {
	err := fetchCoreConfig(EnvDev, APP_GROUP)
	if err != nil { //如果 获取成功
		log.Fatal(err)
		return false
	}
	return true
}

// -----------------------------------------------------------
var CoreConfig map[string]interface{}

var CoreConfigInited bool = false

var CoreConfigLocker *sync.Mutex

//从nacos获取系统级配置
func fetchCoreConfig(conn *Conn, group string) error {
	//初始化配置
	CoreConfig = make(map[string]interface{})
	//读取环境配置
	if conn == nil { //判断是否获取成功
		return errors.New("Nacos env config not found")
	}

	if group == "" { //判断是否获取成功
		return errors.New("Nacos group id not found")
	}

	//CoreConfigLocker.Lock()
	//defer CoreConfigLocker.Unlock() //定义延时 执行完之后 解锁

	if CoreConfigInited { //如果配置初始化过 直接返回
		return nil
	}

	configClient, err := makeConfig(conn)
	if err != nil {
		return errors.New("Nacos env config not found")
	}

	content, err := configClient.GetConfig(vo.ConfigParam{ //获取配置文件
		DataId: DataId,
		Group:  group,
	})
	if err != nil {
		return err
	}

	//把数据转换成byte类型 并且生成切片
	yaml.Unmarshal([]byte(content), CoreConfig) //将数据写入系统配置变量
	fmt.Println(CoreConfig)

	CoreConfigInited = true //代表初始化成功
	return nil
}
