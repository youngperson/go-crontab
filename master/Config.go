package master

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiPort         int      `json:"apiPort"`
	ApiReadTimeout  int      `json:"apiReadTimeout"`
	ApiWriteTimeout int      `json:"apiWriteTimeout"`
	EtcdEndpoints   []string `json:"etcdEndpoints"`
	EtcdDiatimeout  int      `json:"etcdDiatimeout"`
	Webroot         string   `json:"webroot"`
}

var (
	// 单例对象
	G_config *Config
)

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)
	// 1.把配置文件读进来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// 2.做json反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3.赋值单例
	G_config = &conf
	return
}
