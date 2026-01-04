package db

import (
	"database/sql"
	"fmt"
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
		db_name := os.Getenv("DBNAME")
		DB = connectDB(safe_pwd, db_name)
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return DB
}

func connectDB(pwd string, db_name string) *sql.DB {

	cfg := mysql.NewConfig()
	cfg.User = "svc_rw_budget"
	cfg.Passwd = pwd
	cfg.Net = "tcp"
	cfg.Addr = "192.168.40.101:3307"
	cfg.DBName = db_name

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
