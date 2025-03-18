package user


type UserRepositoryI interface {
	// GetUserCollection(ctx context.Context) ([]User, error)

	GetUserItem(userId int) (*User, error)

	// UpdateUserItem(ctx context.Context)

	// DeleteUserItem(cx context.Context)
}
