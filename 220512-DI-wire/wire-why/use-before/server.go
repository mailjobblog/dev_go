package main

type Config struct {
	DbSource string
}

func NewConfig() *Config {
	return &Config{
		DbSource: "root:root@tcp(127.0.0.1:3306)/test_db",
	}
}

type DB struct {
	table string
}

func NewDB(cfg *Config) *DB {
	// TODO 建立mysql连接资源
	return &DB{table: "test_table"}
}

func (db *DB) Find() string {
	return "db info string"
}
