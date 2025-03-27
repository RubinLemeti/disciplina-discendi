package customerr

import "errors"

var ErrUsernameNotUnique = errors.New("username already exists")

var ErrUserNotFound = errors.New("user with given ID not found")
