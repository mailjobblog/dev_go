package db

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"wire-example/internal/server/config"
)

var Provider = wire.NewSet(NewDb)

func NewDb(ctx context.Context, cfg *config.Config) (db *sql.DB, cleanup func(), err error) {
	db, err = sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, func() {
		db.Close()
	}, nil
}
