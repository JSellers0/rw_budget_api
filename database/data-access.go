package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"rw_budget_api/config"
)

var DB *sql.DB

func GetDB() *sql.DB {
	if DB == nil {
		log.Printf("Connecting to Database %s ....", config.DbName)
		DB = connectDB()
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Print("Connected!")
	return DB
}

func connectDB() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = config.DbUser
	cfg.Passwd = config.DbPswd
	cfg.Net = "tcp"
	cfg.Addr = config.DbHost + ":" + config.DbPort
	cfg.DBName = config.DbName

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
