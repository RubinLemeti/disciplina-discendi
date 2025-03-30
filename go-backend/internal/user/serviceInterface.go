package user

type UserServiceI interface {
	GetUserList(limit int, offset int) (*int, []*User, error)

	GetUserItem(userId int) (*User, error)

	AddUserItem(user AddUserItemModel) (*int, error)

	UpdateUserItem(userId int, userBody interface{}) (*int, error)

	DeleteUserItem(userId int) (*int, error)
}
