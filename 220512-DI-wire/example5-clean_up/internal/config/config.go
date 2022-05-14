package config

import (
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"os"
)

var Provider = wire.NewSet(New) // 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象,也就是Config对象

type Config struct {
	Database database `json:"database"`
}

type database struct {
	Dsn string `json:"dsn"`
}

func New() (*Config, func(), error) {
	fp, err := os.Open("../config/app.json")
	if err != nil {
		return nil, func() {}, err
	}
	// defer fp.Close() // 这里注释了，交由外部处理
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, func() {
			fp.Close()
		}, err
	}
	return &cfg, func() {
		fp.Close()
		fmt.Println("app.json 资源句柄关闭成功")
	}, nil
}
