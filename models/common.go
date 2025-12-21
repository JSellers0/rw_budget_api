package models

import (
	"database/sql"
	"net/http"
	da "rw_budget/api/internal"
)

var DB = da.GetDB()

func getErrStatus(err error) (status int) {
	if err == sql.ErrNoRows {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
