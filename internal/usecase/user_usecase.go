package usecase

import (
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
	"github.com/xdevor/ginger/internal/config"
	"github.com/xdevor/ginger/internal/domain"
	"github.com/xdevor/ginger/internal/schema"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (userUsecase *userUsecase) Register(username string, email string, password string) (domain.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	user, err := userUsecase.userRepository.Create(username, email, string(hashedPassword))
	return user, err
}

func (userUsecase *userUsecase) Login(email string, password string) (string, error) {
	user, err := userUsecase.userRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	jwtPreToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(getJwtTTL()) * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    getJwtIssuer(),
		Subject:   strconv.Itoa(int(user.ID)),
	})

	secret := base64.StdEncoding.EncodeToString([]byte(config.Auth.JWT.Secret))
	token, err := jwtPreToken.SignedString([]byte(secret))

	return token, err
}

func getJwtIssuer() string {
	return config.App.Name
}

func getJwtTTL() int {
	ttl, _ := strconv.Atoi(config.Auth.JWT.TTL)
	return ttl
}

func (userUsecase *userUsecase) GetAuthUser(token string) (domain.User, error) {
	var user domain.User

	jwtToken, err := jwt.Parse(token, jwtKeyParseFunc)
	if err != nil {
		return user, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return user, errors.New("unexpected type of jwt claims")
	}

	str, ok := claims["sub"].(string)
	if !ok {
		return user, errors.New("unexpected type of jwt subject")
	}

	id, _ := strconv.Atoi(str)
	user, err = userUsecase.userRepository.GetByID(int64(id))
	if err != nil {
		return user, errors.New("not found user")
	}

	return user, nil
}

func jwtKeyParseFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}
	secret := base64.StdEncoding.EncodeToString([]byte(config.Auth.JWT.Secret))
	return []byte(secret), nil
}

func (userUsecase *userUsecase) CurrentAuthUserID(token string) (int64, error) {
	jwtToken, err := jwt.Parse(token, jwtKeyParseFunc)
	if err != nil {
		return 0, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("unexpected type of jwt claims")
	}

	str, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("unexpected type of jwt subject")
	}

	id, err := strconv.Atoi(str)
	return int64(id), err
}

func (userUsecase *userUsecase) UpdateUser(id int64, updateUser schema.UpdateUser) (domain.User, error) {
	user, err := userUsecase.userRepository.UpdateUser(id, domain.User{
		Email:    updateUser.Email,
		Password: updateUser.Password,
		Username: updateUser.Username,
		Bio:      updateUser.Bio,
		Image:    updateUser.Image,
	})

	return user, err
}
