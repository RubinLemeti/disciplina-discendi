package user

import (
	// "errors"
	"github.com/jinzhu/copier"
	"go-backend/internal/helper/customerr"
	"log/slog"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryI {
	return &UserRepository{db: db}
}

func (ur UserRepository) GetUserList(limit int, offset int) (*int, []*User, error) {
	var userList []*User
	var total int

	rows, err := ur.db.Raw(
		`select id, 
			username, 
			email, 
			password, 
			created_at, 
			updated_at 
			from go_backend.users
			limit ?
			offset ?`, limit, offset).Rows()

	if err != nil {
		slog.Error(err.Error())
		return nil, nil, err
	}
	defer rows.Close()

	row := ur.db.Raw(
		`select count(id) 
		from go_backend.users`).Row()
	row.Scan(&total)

	for rows.Next() {
		var user User
		ur.db.ScanRows(rows, &user)
		userList = append(userList, &user)
	}

	return &total, userList, nil

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

	if result.RowsAffected == 0 {
		return nil, nil //resource does not exist
	}

	return &user, nil
	// return *User
}

func (ur UserRepository) AddUserItem(user AddUserItemModel) (*int, error) {
	tx := ur.db.Begin()
	var userId int

	isUsernameUnique, err := ur.VerifyUsernameIsUnique(tx, user.Username, nil)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	if !*isUsernameUnique {
		return nil, customerr.ErrUsernameNotUnique
	}

	if err := tx.Raw(
		`insert into go_backend.users
		(username, email, password, created_at, updated_at)
		values(?, ?, ?, now(), now())
		RETURNING id`,
		user.Username, user.Email, user.Password).
		Scan(&userId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &userId, nil
}

func (ur UserRepository) UpdateUserItem(userId int, userBody UpdateUserItemModel) (*int, error) {
	tx := ur.db.Begin()

	// get item from db
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

	if result.RowsAffected == 0 {
		return nil, nil //resource does not exist
	}

	// verify unique values
	if userBody.Username != nil {
		isUsernameUnique, err := ur.VerifyUsernameIsUnique(tx, *userBody.Username, &user.Id)
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}

		if !*isUsernameUnique {
			return nil, customerr.ErrUsernameNotUnique
		}
	}

	// update item from db
	copier.Copy(&user, &userBody)

	// update the row in db
	if err := tx.Raw(
		`update go_backend.users
		set (username, email, password, updated_at)
		=(?, ?, ?, now())
		where id=?
		RETURNING id`,
		user.Username, user.Email, user.Password, user.Id).
		Scan(&userId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &userId, nil
}

func (ur UserRepository) DeleteUserItem(userId int) (*int, error) {
	tx := ur.db.Begin()

	result := tx.Exec(
		`delete
		from go_backend.users
		where id = ?`,
		userId)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil //resource does not exist
	}

	return &userId, nil
}

func (ur UserRepository) VerifyUsernameIsUnique(tx *gorm.DB, username string, id *uint) (*bool, error) {
	var exists int

	if id != nil {
		if err := tx.Raw(
			`select 1 
		from go_backend.users 
		where username = ? 
		and id != ? limit 1`, username, id).
			Scan(&exists).Error; err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	} else {
		if err := tx.Raw(
			`select 1 
		from go_backend.users 
		where username = ? 
		limit 1`, username).
			Scan(&exists).Error; err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	}

	isUnique := exists == 0
	return &isUnique, nil
}
