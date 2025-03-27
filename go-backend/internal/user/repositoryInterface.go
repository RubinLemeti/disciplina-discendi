package user


type UserRepositoryI interface {
	// GetUserCollection(ctx context.Context) ([]User, error)

	GetUserItem(userId int) (*User, error)

	AddUserItem(user AddUserItemModel)(*int, error)

	// UpdateUserItem(ctx context.Context)

	// DeleteUserItem(cx context.Context)
}
