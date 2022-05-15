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
