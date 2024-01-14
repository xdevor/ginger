package repository

import (
	"github.com/xdevor/ginger/internal/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (userRepository *userRepository) Create(username string, email string, password string) (domain.User, error) {
	user := domain.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	result := userRepository.db.Create(&user)

	return user, result.Error
}
func (userRepository *userRepository) GetByID(id int64) (domain.User, error) {
	var user domain.User

	result := userRepository.db.First(&user, id)

	return user, result.Error
}

func (userRepository *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User

	result := userRepository.db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func (userRepository *userRepository) UpdateUser(id int64, user domain.User) (domain.User, error) {
	result := userRepository.db.Model(&domain.User{ID: id}).Updates(map[string]interface{}{
		"bio":   user.Bio,
		"image": user.Image,
	})

	return user, result.Error
}
