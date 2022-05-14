package data

import (
	"context"
	"database/sql"
	"fmt"
	"wire-example/internal/biz"
)

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
	err := o.Dao.QueryRow("SELECT * FROM order WHERE id=?", id).Scan(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepo) Create(ctx context.Context, order *biz.Order) (int64, error) {
	sqlStr := "insert into order(name, type, price) values (?,?,?)"
	ret, err := o.Dao.Exec(sqlStr, order.Name, order.Type, order.Price)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0, err
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return 0, err
	}
	//log.Printf("insert success, the id is %d.\n", theID)
	return theID, nil
}
