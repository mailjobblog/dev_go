//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/internal/config"
)

func InitApp() (*App, func(), error) {
	wire.Build(config.Provider, NewApp)
	return &App{}, func() {}, nil
}
