package user

type UserService struct {
	repository UserRepositoryI
}

func NewUserService(repo UserRepositoryI) UserServiceI {
	return &UserService{repository: repo}
}

func (us UserService) GetUserItem(userId int) (*User, error) {
	return us.repository.GetUserItem(userId)
}

func (us UserService) AddUserItem(user AddUserItemModel) (*int, error) {
	return us.repository.AddUserItem(user)
}
