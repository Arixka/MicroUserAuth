package api

import (
	"github.com/gin-gonic/gin"
	"github.com/microservices/microUserAuth/internal/interface/handlers"
	"github.com/microservices/microUserAuth/internal/usecase/service"
)

func newRouter(userService service.UserService, authService service.AuthService) *gin.Engine{
	router := gin.Default()
	
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	router.POST("/users", userHandler.Register)
	router.POST("/login", authHandler.Login)
	return router
}