package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xdevor/ginger/internal/delivery/http/middleware"
	"github.com/xdevor/ginger/internal/domain"
	"github.com/xdevor/ginger/internal/schema"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

var validate *validator.Validate = validator.New()

func RegisterUserHandler(
	r *gin.Engine,
	userUsecase domain.UserUsecase,
	authMid middleware.AuthMiddleware,
) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}

	api := r.Group("/api")
	{
		api.POST("/users", handler.Register)
		api.POST("/users/login", handler.Login)

		auth := api.Use(authMid.Auth)
		{
			auth.GET("/user", authMid.Auth, handler.GetUser)
			auth.PUT("/user", authMid.Auth, handler.Update)
		}
	}
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	requestPayload := map[string]schema.NewUser{"user": {}}
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := requestPayload["user"]
	if err := validate.Struct(newUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user, err := userHandler.userUsecase.Register(newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.Schema(),
	})
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	requestPayload := map[string]schema.LoginUser{"user": {}}
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginUser := requestPayload["user"]
	if err := validate.Struct(loginUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	token, err := userHandler.userUsecase.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": schema.User{Token: token},
	})
}

func (userHandler *UserHandler) GetUser(c *gin.Context) {
	authUser := c.MustGet("auth_user").(domain.User)

	c.JSON(http.StatusOK, gin.H{
		"user": authUser.Schema(),
	})
}

func (userHandler *UserHandler) Update(c *gin.Context) {
	authUser := c.MustGet("auth_user").(domain.User)

	requestPayload := map[string]schema.UpdateUser{"user": {}}
	if err := c.ShouldBind(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser := requestPayload["user"]
	if err := validate.Struct(updateUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := userHandler.userUsecase.UpdateUser(authUser.ID, updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": updatedUser.Schema(),
	})
}
