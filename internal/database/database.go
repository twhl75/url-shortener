package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	cfg := mysql.Config{
		Net:    "tcp",
		Addr:   "mysql:3306",
		DBName: "urlDB",
		User:   "root",
		Passwd: "password",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db
}
