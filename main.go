package main

import (
	"Atividade5-DB/database"
	_ "Atividade5-DB/docs"
	"fmt"
	"log"
	"os"

	"Atividade5-DB/util"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DBHOST")
	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbPort := os.Getenv("DBPORT")
	port := os.Getenv("PORT")

	connStr :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName)

	db := database.NewDatabase(connStr)
	defer db.Conn.Close()

	util.InjectDependencies(db, port)
}
