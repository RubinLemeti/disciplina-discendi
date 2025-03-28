package user

type AddUserItemModel struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserIdModel struct {
	Id int `json:"id" validate:"gte=0"`
}

type GetUserItemListModel struct{
	Data []UserIdModel
	Metadata int
}