package user

type UserServiceI interface {
	GetUserItem(userId int) (*User, error)
}
