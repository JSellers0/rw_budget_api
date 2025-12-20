package models

import (
	"database/sql"
	"log"
	"net/http"
	da "rw_budget/api/internal"

	"github.com/gin-gonic/gin"
)

var DB = da.GetDB()

type Account struct {
	ID       int    `json:"accountid"`
	Name     string `json:"account_name"`
	Type     string `json:"account_type"`
	Features string `json:"reward_features"`
	PmtDate  string `json:"payment_day"`
	StmtDate string `json:"statement_day"`
}

func GetAccounts(c *gin.Context) {
	base_query := "SELECT accountid, account_name, account_type, rewards_features, payment_day, statement_day\n"
	base_query += "FROM account\n"

	if c.Query("name") != "" {
		getByName(c, base_query)
	} else if c.Query("id") != "" {
		getByID(c, base_query)
	} else {
		getAll(c, base_query)
	}
}

func getAll(c *gin.Context, base_query string) {
	accounts := []Account{}
	results, err := DB.Query(base_query)
	if err != nil {
		log.Print(err.Error())
	}
	for results.Next() {
		var account Account
		err = results.Scan(
			&account.ID,
			&account.Name,
			&account.Type,
			&account.Features,
			&account.PmtDate,
			&account.StmtDate,
		)
		if err != nil {
			log.Print(err.Error())
		}
		accounts = append(accounts, account)
	}
	c.IndentedJSON(http.StatusOK, accounts)
}

func getByName(c *gin.Context, base_query string) {
	base_query += "WHERE account_name = ?\n;"
	var account Account
	if err := DB.QueryRow(
		base_query, c.Query("name")).Scan(
		&account.ID,
		&account.Name,
		&account.Type,
		&account.Features,
		&account.PmtDate,
		&account.StmtDate,
	); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, account)
		}
	}
	c.IndentedJSON(http.StatusOK, account)
}

func getByID(c *gin.Context, base_query string) {
	base_query += "WHERE accountid = ?\n;"
	var account Account
	if err := DB.QueryRow(
		base_query, c.Query("id")).Scan(
		&account.ID,
		&account.Name,
		&account.Type,
		&account.Features,
		&account.PmtDate,
		&account.StmtDate,
	); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, account)
		}
	}
	c.IndentedJSON(http.StatusOK, account)
}

func createAccount(deets Account) int {
	return 0
}

func updateAccount(deets Account) int {
	return 0
}

func deleteAccount(id int) int {
	isAuth := true
	if isAuth {
		return 0
	}
	return 1
}
