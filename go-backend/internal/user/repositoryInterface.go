package user


type UserRepositoryI interface {
	GetUserList(limit int, offset int) (*int, []*User, error)

	GetUserItem(userId int) (*User, error)

	AddUserItem(user AddUserItemModel)(*int, error)

	UpdateUserItem(userId int, userBody UpdateUserItemModel) (*int, error)

	DeleteUserItem(userId int) (*int, error)
}
