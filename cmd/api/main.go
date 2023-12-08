package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/microservices/microUserAuth/internal/infrastructure/database"
)

func main() {

    db, err := database.Connect()
    if err != nil {
        log.Fatalf("Could not establish connection to the database: %v", err)
    }
    defer db.Close()

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Bienvenido a microUserAuth",
        })
    })
    
    //Iniciar el servidor, si error es nil todo bien, sino es nil salta el log
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}