package user

type UserService struct {
	repository UserRepositoryI
}

func NewUserService(repo UserRepositoryI) UserServiceI {
	return &UserService{repository: repo}
}

func (us UserService) GetUserList(limit int, offset int) (*int, []*User, error) {
	return us.repository.GetUserList(limit, offset)
}

func (us UserService) GetUserItem(userId int) (*User, error) {
	return us.repository.GetUserItem(userId)
}

func (us UserService) AddUserItem(user AddUserItemModel) (*int, error) {
	return us.repository.AddUserItem(user)
}

func (us UserService) UpdateUserItem(userId int, userBody UpdateUserItemModel) (*int, error) {
	return us.repository.UpdateUserItem(userId, userBody)
}

func (us UserService) DeleteUserItem(userId int) (*int, error) { 
	return us.repository.DeleteUserItem(userId)
 }
