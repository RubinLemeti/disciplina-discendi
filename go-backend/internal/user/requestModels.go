package user

type AddUserItemModel struct {
	Username string `json:"username" validate:"required,min=3,excludesall= "`
	Email    string `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"required,min=5,alphanum,excludesall= "`
}

type UpdateUserItemModel struct {
	Username *string `json:"username" validate:"omitempty,alpha,min=3,excludesall= "`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password" validate:"omitempty,min=5,alphanum,excludesall= "`
}

type UserIdModel struct {
	Id int `json:"id" validate:"gte=0"`
}

type GetUserListQueryParams struct {
	Limit  int `json:"limit" validate:"gte=0"`
	Offset int `json:"offset" validate:"gte=-1"`
}
