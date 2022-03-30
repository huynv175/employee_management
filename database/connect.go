package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// connect to db

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  os.Getenv("DBNET"),
		Addr:                 os.Getenv("DBADDRESS"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}
	fmt.Printf("config: %+v\n", cfg)
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Ping status: OK")
	}
	DB = db
}
