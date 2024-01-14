package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xdevor/ginger/internal/domain"
)

type AuthMiddleware struct {
	UserUsecase domain.UserUsecase
}

func NewAuthMiddleware(userUsecase domain.UserUsecase) AuthMiddleware {
	return AuthMiddleware{
		UserUsecase: userUsecase,
	}
}

func (authMiddleware AuthMiddleware) Auth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	authUser, err := authMiddleware.UserUsecase.GetAuthUser(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("auth_user", authUser)
	c.Next()
}
