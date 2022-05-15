package db

import "database/sql"

// IDao 为接口，Dao 为实现，Dao 依赖于 IDao

// 接口声明
type IDao interface {
	Version() (string, error)
}

// 默认实现
type Dao struct { // 默认实现
	db *sql.DB
}

// 生成dao对象的方法
func NewDao(db *sql.DB) IDao {
	return &Dao{
		db: db,
	}
}

// 在 Dao 中实现 IDao 的接口
func (d *Dao) Version() (string, error) {
	var version string
	row := d.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		return "", err
	}
	return version, nil
}
