package models

import (
	"database/sql"
	"net/http"
	db "rw_budget/api/internal"
)

var DB = db.GetDB()

func getErrStatus(err error) (status int) {
	if err == sql.ErrNoRows {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
