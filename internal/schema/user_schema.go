package schema

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type NewUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	Username string `json:"username" validate:"required,min=6,max=32"`
}

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=6,max=32"`
	Token    string `json:"token" validate:"required"`
	Bio      string `json:"bio" validate:"required"`
	Image    string `json:"image" validate:"required"`
}

type UpdateUser struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6,max=32"`
	Username string `json:"username" validate:"min=6,max=32"`
	Bio      string `json:"bio" validate:"min=1,max=300"`
	Image    string `json:"image" validate:"filepath"`
}
