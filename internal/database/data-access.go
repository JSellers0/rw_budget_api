package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func GetDB() *sql.DB {
	if DB == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		safe_pwd := os.Getenv("MSQLPWD")
		DB = connectDB(safe_pwd)
	}
	return DB
}

func connectDB(pwd string) *sql.DB {

	cfg := mysql.NewConfig()
	cfg.User = "svc_rw_budget"
	cfg.Passwd = pwd
	cfg.Net = "tcp"
	cfg.Addr = "192.168.40.101:3307"
	cfg.DBName = "rw_budget_dev"

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
