package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

// 客户端连接
type Conn struct {
	Addr         string
	SlaveAddr    string
	Namespace    string
	Path         string
	Port         uint64
	CacheDir     string
	LogsDir      string
	BeatInterval int64
	TimeoutMs    uint64
}

//生成配置
func makeConfig(conn *Conn) (config_client.IConfigClient, error) {
	clientConfig := constant.ClientConfig{
		TimeoutMs:    conn.TimeoutMs,
		BeatInterval: conn.BeatInterval,
		CacheDir:     conn.CacheDir,
		LogDir:       conn.LogsDir,
		// 我们可以创建多个具有不同名称空间ID的客户端，以支持多个名称空间。如果名称空间是公共的，请在此处填写空白字符串
		NamespaceId: conn.Namespace,
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      conn.Addr,
			ContextPath: conn.Path,
			Port:        conn.Port,
		},
		{
			IpAddr:      conn.SlaveAddr,
			ContextPath: conn.Path,
			Port:        conn.Port,
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{ //初始化Nacos客户端
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	return configClient, err
}
