package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"wire-example/internal/server/config"
)

var Provider = wire.NewSet(NewDb)

func NewDb(cfg *config.Config) (db *sql.DB, cleanup func(), err error) {
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
