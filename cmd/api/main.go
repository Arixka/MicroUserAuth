package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()


    //Iniciar el servidor, si error es nil todo bien, sino es nil salta el log
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}