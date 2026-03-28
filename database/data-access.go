package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"rw_budget_api/config"
)

var DB *sql.DB

func GetDB() error {
	if DB != nil {
		return nil // Already Connected
	}
	log.Printf("Connecting to Database %s ...", config.DbName)
	var err error
	DB, err = sql.Open("mysql", getDbDSN())
	if err != nil {
		return err
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		return pingErr
	}

	log.Print("Connected!")
	return nil
}

func getDbDSN() string {
	cfg := mysql.NewConfig()
	cfg.User = config.DbUser
	cfg.Passwd = config.DbPswd
	cfg.Net = "tcp"
	cfg.Addr = config.DbHost + ":" + config.DbPort
	cfg.DBName = config.DbName

	return cfg.FormatDSN()
}
