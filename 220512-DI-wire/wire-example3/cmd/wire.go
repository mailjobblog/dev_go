

package main

import (
	"github.com/google/wire"
	"wire-example/internal/config"
	"wire-example/internal/db"
	"wire-example/internal/services"
)

func InitApp() (*App, error)  *services.UserService {
	wire.Build(config.Provider, db.Provider, NewApp)
	return nil
}
