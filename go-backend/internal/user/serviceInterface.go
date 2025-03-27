package user

type UserServiceI interface {
	GetUserItem(userId int) (*User, error)

	AddUserItem(user AddUserItemModel)(*int, error)
}
