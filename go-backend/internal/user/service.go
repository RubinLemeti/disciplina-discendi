package user

type UserService struct {
	repository UserRepositoryI
}

func NewUserService(repo UserRepositoryI) UserServiceI {
	return &UserService{repository: repo}
}

func (us UserService) GetUserList() ([]*User, error) {
	return us.repository.GetUserList()
}

func (us UserService) GetUserItem(userId int) (*User, error) {
	return us.repository.GetUserItem(userId)
}

func (us UserService) AddUserItem(user AddUserItemModel) (*int, error) {
	return us.repository.AddUserItem(user)
}

func (us UserService) UpdateUserItem(userId int, userBody interface{}) (*int, error) {return nil, nil}

func (us UserService) DeleteUserItem(userId int) (*int, error) { return nil, nil }
