package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/microservices/microUserAuth/internal/infrastructure/database"
	"github.com/microservices/microUserAuth/internal/infrastructure/middlewares"
	repositoryimpl "github.com/microservices/microUserAuth/internal/infrastructure/repositoy_impl"
	"github.com/microservices/microUserAuth/internal/interface/handlers"
	"github.com/microservices/microUserAuth/internal/usecase/service"
)

func main() {

	envFile := ".env.local"
	if os.Getenv("DOCKER_ENV") == "true" {
		envFile = ".env.docker"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not establish connection to the database: %v", err)
	}

	userRepo := repositoryimpl.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	router := gin.Default()
	router.Use(middlewares.CORS())

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	router.POST("/users", userHandler.Register)
	router.POST("/login", authHandler.Login)

	log.Printf("************************************************************\n********************************************************************************")
	//Iniciar el servidor, si error es nil todo bien, sino es nil salta el log
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
