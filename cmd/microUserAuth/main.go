package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Bienvenido a microUserAuth",
        })
    })

    r.Run() // Por defecto, escucha en el puerto 8080
}