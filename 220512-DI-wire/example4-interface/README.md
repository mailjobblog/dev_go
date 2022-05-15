# Wire 接口绑定

在面向接口编程中，代码依赖的往往是接口，而不是具体的struct。

在此 demo 中，我们基于 `example1` 中的 demo 进行修改，将数据库查询逻辑封装为借口，引入交由 main 方法调用。

internal/db/dao.go

```go
package db

import "database/sql"

// 接口声明
type IDao interface {
	Version() (string, error)
}

// 默认实现
type Dao struct { // 默认实现
	db *sql.DB
}

// 生成dao对象的方法
func NewDao(db *sql.DB) *Dao {
	return &Dao{
		db: db,
	}
}

// 在 Dao 中实现 IDao 的借口
func (d *Dao) Version() (string, error) {
	var version string
	row := d.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		return "", err
	}
	return version, nil
}
```

internal/db/db.go

```go
package db

// ... ...

// 这里我们加入了 Dao，并且绑定了 IDao 和 Dao
var Provider = wire.NewSet(NewDb, NewDao, wire.Bind(new(IDao), new(*Dao))) // 这里将接口和实现进行绑定

func NewDb(cfg *config.Config) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, nil
}
```

cmd/main.go

```go
package main

// ... ...

type App struct {
	dao db.IDao
}

func NewApp(i db.IDao) *App {
	return &App{
		dao: i,
	}
}

func main() {
	app, _ := InitApp()
	version, _ := app.dao.Version()
	fmt.Println(version)
}
```

**代码分析**

这里最重要的是这一段代码，我们来分析一下

```go
var Provider = wire.NewSet(NewDb, NewDao, wire.Bind(new(IDao), new(*Dao))) 
```

在 NewDao 中注入了 NewDb 方法。

这里的 NewDao 返回自定义的 Dao 结构体，并且把 NewDb 注入的 DB 保存到 Dao 结构体中，提供给 Dao 实现的方法使用。

wire.Bind将 IDao 借口和 Dao 实现进行绑定。

最终调用 IDao 的接口，即可调用到 Dao 中实现的方法。

