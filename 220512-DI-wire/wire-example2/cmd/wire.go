//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/internal/config"
	"wire-example/internal/db"
)

func InitApp() (*App, error) {
	wire.Build(config.Provider, db.Provider, NewApp)
	return &App{}, nil
}
