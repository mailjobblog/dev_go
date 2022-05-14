//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"wire-example/internal/data"
	"wire-example/internal/server/config"
	"wire-example/internal/server/db"
)

func InitApp(ctx context.Context) (*App, func(), error) {
	panic(wire.Build(config.Provider, db.Provider, data.OrderSet, NewApp))
}
