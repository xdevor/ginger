package domain

import "github.com/xdevor/ginger/internal/schema"

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at" gorm:"default:null"`
	UpdatedAt string `json:"updated_at" gorm:"default:null"`
}

type UserUsecase interface {
	Register(username string, email string, password string) (User, error)
	Login(email string, password string) (token string, err error)
	GetAuthUser(token string) (User, error)
	UpdateUser(int64, schema.UpdateUser) (User, error)
}

type UserRepository interface {
	Create(username string, email string, password string) (User, error)
	GetByID(id int64) (User, error)
	GetByEmail(email string) (User, error)
	UpdateUser(id int64, user User) (User, error)
}

func (user User) Schema() schema.User {
	return schema.User{
		Email:    user.Email,
		Token:    user.Token,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
	}
}
