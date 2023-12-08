package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//usamos una importacion anonima porque al no usarlo directamente el compilador lo considera no necesario y no lo importa
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error){
	//obtenemos las variables de entorno
	host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

	// Creamos la cadena de conexión
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
	
	// conectamos con la base de datos
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, fmt.Errorf("Error opening database: %v", err)
    }

    // Verificamos la conexion
    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("Error connecting to database: %v", err)
    }

    log.Println("Successfully connected to database")
    return db, nil
}