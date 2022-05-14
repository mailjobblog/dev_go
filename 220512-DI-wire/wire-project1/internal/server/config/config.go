package config

import (
	"encoding/json"
	"github.com/google/wire"
	"os"
)

var Provider = wire.NewSet(New)

type Config struct {
	Database database `json:"database"`
}

type database struct {
	Dsn string `json:"dsn"`
}

func New() (*Config, func(), error) {
	fp, err := os.Open("../../config/app.json")
	if err != nil {
		return nil, func() {}, err
	}
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, func() {}, err
	}
	return &cfg, func() {
		fp.Close()
	}, nil
}
