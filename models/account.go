package models

import (
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
	var status int
	var records []Account

	base_query := "SELECT accountid, account_name, account_type, rewards_features, payment_day, statement_day\n"
	base_query += "FROM account\n"

	if c.Query("id") != "" {
		status, records = getAccountByID(base_query, c.Query("id"))
	} else if c.Query("name") != "" {
		status, records = getAccountByName(base_query, c.Query("name"))
	} else {
		status, records = getAllAccounts(base_query)
	}
	c.IndentedJSON(status, records)
}

func getAccountByID(base_query string, id string) (status int, record []Account) {
	base_query += "WHERE accountid = ?\n;"
	var account Account
	if err := DB.QueryRow(
		base_query, id).Scan(
		&account.ID,
		&account.Name,
		&account.Type,
		&account.Features,
		&account.PmtDate,
		&account.StmtDate,
	); err != nil {
		log.Print(err.Error())
		return handleSqlErr(err), []Account{}
	}
	return http.StatusOK, []Account{account}
}

func getAccountByName(base_query string, name string) (status int, record []Account) {
	base_query += "WHERE account_name = ?\n;"
	var account Account
	if err := DB.QueryRow(
		base_query, name).Scan(
		&account.ID,
		&account.Name,
		&account.Type,
		&account.Features,
		&account.PmtDate,
		&account.StmtDate,
	); err != nil {
		return handleSqlErr(err), []Account{}
	}
	return http.StatusOK, []Account{account}
}

func getAllAccounts(base_query string) (status int, records []Account) {
	accounts := []Account{}
	results, err := DB.Query(base_query)
	if err != nil {
		return handleSqlErr(err), accounts
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
			return http.StatusInternalServerError, []Account{}
		}
		accounts = append(accounts, account)
	}
	return http.StatusOK, accounts
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
