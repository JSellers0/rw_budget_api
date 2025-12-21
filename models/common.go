package models

import (
	"database/sql"
	"net/http"
)

func handleSqlErr(err error) (status int) {
	if err == sql.ErrNoRows {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
