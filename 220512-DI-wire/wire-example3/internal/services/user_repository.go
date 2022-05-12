package services

// UserRepository 存放User对象的数据仓库接口,eg: mysql,restful api ....
type UserRepository interface {
	// GetUserByID 根据ID获取User, 如果找不到User返回对应错误信息
	GetUserByID(id int) (*User, error)
}
