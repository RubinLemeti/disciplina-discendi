package user

type AddUserItemModel struct {
	Email    string `json:"email" validate:"required,email"`
	UserName string `json:"userName" validate:"required"`
}
