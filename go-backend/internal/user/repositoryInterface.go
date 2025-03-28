package user


type UserRepositoryI interface {
	GetUserList() ([]*User, error)

	GetUserItem(userId int) (*User, error)

	AddUserItem(user AddUserItemModel)(*int, error)

	UpdateUserItem(userId int, userBody interface{}) (*int, error)

	DeleteUserItem(userId int) (*int, error)
}
