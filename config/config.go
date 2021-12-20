package config

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB


type config struct{
    User string `json:"user"`
    Password string `json:"passwd"`
    Net string `json:"net"`
    Address string `json:"addr"`
    DB string `json:"dbname"`
    AllowNativePassword  bool `json:"ANP"`
	Database string `json:"Database"`
}

func Connect() *sql.DB {
	b, err := ioutil.ReadFile("./config/config.json")
	
	if err != nil {
	 log.Print(err)
	}
	var cfgJson config;
	err = json.Unmarshal(b,&cfgJson)
	if err != nil {
		panic(err)
	}
	
	cfg := mysql.Config{
		User:                 cfgJson.User,
		Passwd:               cfgJson.Password,
		Net:                  cfgJson.Net,
		Addr:                 cfgJson.Address,
		DBName:               cfgJson.DB,
		AllowNativePasswords: cfgJson.AllowNativePassword} 
	//get the database
	db, err = sql.Open( "mysql", cfg.FormatDSN())
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

func Getdb() *sql.DB{
	if(db != nil){
		return db
	}
	Connect();
	return db;
}