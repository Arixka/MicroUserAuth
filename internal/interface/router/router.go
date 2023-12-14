package api

import (
	"github.com/gin-gonic/gin"
	"github.com/microservices/microUserAuth/internal/interface/handlers"
	"github.com/microservices/microUserAuth/internal/usecase/service"
)

func newRouter(userService service.UserService) *gin.Engine{
	router := gin.Default()
	
	userHandler := handlers.NewUserHandler(userService)
	router.POST("/users", userHandler.Register)

	return router
}