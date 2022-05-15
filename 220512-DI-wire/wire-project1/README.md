# Wire 工程化实践

在 wire 工程化中， 我主要参考了以下项目对于该 demo 进行实现，为了便于观看，部分内容做了精简。

**项目布局参考：**

- layout-project：https://github.com/golang-standards/project-layout
- Kratos框架：https://github.com/go-kratos/kratos

**wire实践参考：**

- 基于Kratos框架的微服务项目：https://github.com/go-kratos/beer-shop

## Demo详解

该demo项目布局

```text
_
│  go.mod
│  go.sum
├─cmd
│      main.go
│      wire.go
│      wire_gen.go
├─config
│      app.json
└─internal
    ├─biz
    │      order.go
    ├─data
    │      data.go
    │      order.go
    └─server
        ├─config
        │      config.go
        └─db
                mysql.go
```

### internal 目录详解

该目录存放所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用

**internal/biz**

业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。

**internal/data**

业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。

**internal/server**

config，db，http，grpc 实例的创建和配置

### 工程化实践

在该演示示例中，需求是要求 订单查询 和 订单创建，所以我们现需要创建对应 order 表和测试数据。

```bash
# 创建数据库和表

# 添加测试数据
```

internal/biz/order.go

在领域层，我定义了关于 order 表的映射结构体，并且定义了业务逻辑的接口

```go
package biz

import "context"

type Order struct {
	Id    int64
	Name  string
	Price float64
}

type OrderRepo interface {
	// Find 根据ID查询一条记录
	Find(context.Context, int64) (*Order, error)
	// Create 创建一条记录
	Create(context.Context, *Order) (int64, error)
}
```

internal/data/data.go

在数据持久层的 data.go 文件中，我将 mysql 驱动注入到 Data 中，便于在接口实现方法中，调用 db 驱动。

```go
package data

// ... ...

var OrderSet = wire.NewSet(NewData, NewOrderRepo)

type Data struct {
	Mysql *sql.DB
}

func NewData(mysql *sql.DB) (*Data, func(), error) {
	return &Data{
		Mysql: mysql,
	}, func() {}, nil
}
```

internal/data/order.go

```go
package data

// ... ...

// 要求 OrderRepo 必须实现 biz.OrderRepo 所有接口
var _ biz.OrderRepo = (*OrderRepo)(nil)

type OrderRepo struct {
	Dao *sql.DB
}

// NewOrderRepo
// 实现 接口(biz.OrderRepo) 与 实现(OrderRepo) 的绑定关系
// 此方法定义的返回是 接口 实际返回的是 具体对象
// 该 具体对象 已经实现了 接口 的所有方法
// 这样调用 biz.OrderRepo 中的 [方法] 即可调用到 OrderRepo 的 [方法]
func NewOrderRepo(data *Data) (biz.OrderRepo, func(), error) {
	return &OrderRepo{
		Dao: data.Mysql,
	}, func() {}, nil
}

func (o *OrderRepo) Find(ctx context.Context, id int64) (*biz.Order, error) {
	var order biz.Order
	err := o.Dao.QueryRow("SELECT * FROM `order` WHERE id=?", id).Scan(&order.Id, &order.Name, &order.Price)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepo) Create(ctx context.Context, order *biz.Order) (int64, error) {
	sqlStr := " INSERT INTO `order`(`name`, `price`) values (?,?)"
	ret, err := o.Dao.Exec(sqlStr, order.Name, order.Price)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0, err
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return 0, err
	}
	return theID, nil
}
```

cmd/main.go

```go
package main

// ... ...

type App struct {
	conf      *config.Config
	db        *sql.DB
	OrderRepo biz.OrderRepo
}

func NewApp(cfg *config.Config, db *sql.DB, orderRepo biz.OrderRepo) *App {
	return &App{
		conf:      cfg,
		db:        db,
		OrderRepo: orderRepo,
	}
}

func main() {
	ctx := context.Background()
	app, cleanup, err := InitApp(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// 测试 db 查询
	order, err := app.OrderRepo.Find(ctx, 1)
	if err != nil {
		fmt.Println("order find error:", err)
		return
	}
	fmt.Printf("查询成功， %+v", order)
	fmt.Println()

	// 测试 db 插入
	var o = biz.Order{
		Name:  "测试商品",
		Price: 9.9,
	}
	ID, err := app.OrderRepo.Create(ctx, &o)
	if err != nil {
		return
	}
	fmt.Printf("插入成功，id = %d", ID)
}

```


> 什么是接口？  
> 接口的特性是golang支持鸭子类型的基础，即“如果它走起来像鸭子，叫起来像鸭子（实现了接口要的方法），它就是一只鸭子（可以被赋值给接口的值）”

**注意：**

NewOrderRepo 方法定义返回的是 biz.OrderRepo 接口，实际返回的是 OrderRepo 实现。

在 main 中 App OrderRepo 定义的类型是  biz.OrderRepo 所以和上文是相呼应的。

比较值得注意的事，上面两个  biz.OrderRepo 都没有写 * ，是因为接口本身就是指针，所以不需要添加 *。



