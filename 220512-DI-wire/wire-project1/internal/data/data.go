package data

import (
	"database/sql"
	"github.com/google/wire"
)

var OrderSet = wire.NewSet(NewData, NewOrderRepo)

type Data struct {
	Mysql *sql.DB
}

func NewData(mysql *sql.DB) (*Data, func(), error) {
	return &Data{
		Mysql: mysql,
	}, func() {}, nil
}
