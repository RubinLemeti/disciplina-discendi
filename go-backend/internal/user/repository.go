package user

import (
	// "errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryI {
	return &UserRepository{db: db}
}

func (ur UserRepository) GetUserItem(userId int) (*User, error) {
	var user User
	result := ur.db.Raw(
		`select id, 
			username, 
			email, 
			password, 
			created_at, 
			updated_at 
			from go_backend.users
			where id=?`, 
		userId).Scan(&user)
	

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0{
		return nil, nil //resource does not exist
	}

	return &user, nil
	// return *User
}



