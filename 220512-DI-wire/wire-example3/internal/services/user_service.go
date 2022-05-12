package services

// UserService 关于 user 的一系列的仓库
type UserService struct {
	userRepo UserRepository // <-- UserService依赖UserRepository接口
}

// NewUserService *UserService构造函数
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// UserExist 判断指定ID的用户是否存在
func (u *UserService) UserExist(id int) bool {
	_, err := u.userRepo.GetUserByID(id)
	return err == nil
}
