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

type String2 string

func New(s string, s2 String2) (*Config, error) {

	fmt.Println("print wire value：" + s)
	fmt.Println("print wire value：" + s2)

	fp, err := os.Open("../config/app.json")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
