package db

import "database/sql"

type IDao interface { // 接口声明
	Version() (string, error)
}

type Dao struct { // 默认实现
	db *sql.DB
}

func NewDao(db *sql.DB) *Dao { // 生成dao对象的方法
	return &Dao{db: db}
}

func (d *Dao) Version() (string, error) {
	var version string
	row := d.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		return "", err
	}
	return version, nil
}
