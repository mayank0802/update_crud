package config

import (
	"database/sql"
	"fmt"
	"log"

	//"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() *sql.DB {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "DB",
		AllowNativePasswords: true,
	}

	//get a databasr
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
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
