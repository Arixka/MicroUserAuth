package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/microservices/microUserAuth/internal/infrastructure/database"
	repositoryimpl "github.com/microservices/microUserAuth/internal/infrastructure/repositoy_impl"
	"github.com/microservices/microUserAuth/internal/interface/handlers"
	"github.com/microservices/microUserAuth/internal/usecase/service"
)

func main() {

    db, err := database.Connect()
    if err != nil {
        log.Fatalf("Could not establish connection to the database: %v", err)
    }

    userRepo := repositoryimpl.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	router := gin.Default()

	userHandler := handlers.NewUserHandler(userService)
    
	router.POST("/users", userHandler.Register)

    //Iniciar el servidor, si error es nil todo bien, sino es nil salta el log
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}