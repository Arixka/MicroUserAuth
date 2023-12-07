package api

import (
	"github.com/gin-gonic/gin"
)

func newRouter(userHandler *UserHandler) *gin.Engine{
    router := gin.Default()

	router.POST("/users", userHandler.Register)

	return router
}