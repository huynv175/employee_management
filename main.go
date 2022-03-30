package main

import (
	"employee-management-system/database"
	"employee-management-system/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	database.ConnectDB()
	err := godotenv.Load("./enviroment/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	routers.Router()
}
