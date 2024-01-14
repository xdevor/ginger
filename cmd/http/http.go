package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/xdevor/ginger/internal/client"
	"github.com/xdevor/ginger/internal/config"
	"github.com/xdevor/ginger/internal/delivery/http"
	"github.com/xdevor/ginger/internal/delivery/http/middleware"
	"github.com/xdevor/ginger/internal/repository"
	"github.com/xdevor/ginger/internal/usecase"
)

func main() {
	// Boot DB
	db, err := client.NewPostgres()
	if err != nil {
		panic(err.Error())
	}

	// Boot Repository
	userRepository := repository.NewUserRepository(db)

	// Boot Usecase
	userUsecase := usecase.NewUserUsecase(userRepository)

	// Boot Middleware
	authMid := middleware.NewAuthMiddleware(userUsecase)

	// Boot Delivery
	gin := gin.Default()
	http.RegisterUserHandler(gin, userUsecase, authMid)

	gin.Run(":" + config.App.Port)
}
