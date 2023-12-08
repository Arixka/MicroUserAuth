package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error){
	//obtenemos las variables de entorno
	host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

	// Creamos la cadena de conexión
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
	
	// conectamos con la base de datos usando gorm

    db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
    if err != nil {
        log.Printf("Error opening database connection: %v", err)
        return nil, err
    }

    
    sqlDB, err := db.DB()
    if err != nil {
        log.Printf("Error getting SQL DB from GORM: %v", err)
        return nil, err
    }
    
    // Verificamos la conexion
    err = sqlDB.Ping()
    if err != nil {
        log.Printf("Error on pinging database: %v", err)
        return nil, err
    }

    log.Println("Successfully connected nd pinged database")
    return db, nil
}