package da

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB(pwd string) *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "svc_rw_budget"
	cfg.Passwd = pwd
	cfg.Net = "tcp"
	cfg.Addr = "192.168.40.101:3306"
	cfg.DBName = "rw_budget_dev"

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
