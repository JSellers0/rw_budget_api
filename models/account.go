package models

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	ID       int64  `json:"accountid"`
	Name     string `json:"account_name" binding:"required"`
	Type     string `json:"account_type" binding:"required"`
	Features string `json:"rewards_features"`
	PmtDay   string `json:"payment_day"`
	StmtDay  string `json:"statement_day"`
}

func GetAccounts(c *gin.Context) {
	var records []Account
	var err error
	base_query := getBaseQuery()

	if c.Query("id") != "" {
		records, err = getAccountByID(base_query, c.Query("id"))
	} else if c.Query("name") != "" {
		records, err = getAccountByName(base_query, c.Query("name"))
	} else {
		records, err = getAllAccounts(base_query)
	}

	if err != nil {
		log.Print("Error retrieving accounts.")
		log.Print(err.Error())
		c.IndentedJSON(getErrStatus(err), records)
	}
	c.IndentedJSON(http.StatusOK, records)
}

func getAccountByID(base_query string, id string) (records []Account, err error) {
	base_query += "WHERE accountid = ?\n;"
	var account Account
	if err := DB.QueryRow(
		base_query, id).Scan(
		&account.ID,
		&account.Name,
		&account.Type,
		&account.Features,
		&account.PmtDay,
		&account.StmtDay,
	); err != nil {
		return nil, err
	}
	return []Account{account}, nil
}

func getAccountByName(base_query string, name string) (records []Account, err error) {
	base_query += "WHERE account_name = ?\n;"
	var account Account
	if err := DB.QueryRow(
		base_query, name).Scan(
		&account.ID,
		&account.Name,
		&account.Type,
		&account.Features,
		&account.PmtDay,
		&account.StmtDay,
	); err != nil {
		return nil, err
	}
	return []Account{account}, nil
}

func getAllAccounts(base_query string) (records []Account, err error) {
	accounts := []Account{}
	results, err := DB.Query(base_query)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var account Account
		err = results.Scan(
			&account.ID,
			&account.Name,
			&account.Type,
			&account.Features,
			&account.PmtDay,
			&account.StmtDay,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func CreateAccount(c *gin.Context) {
	new_account, err := bindAccount(c)
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	account, get_err := getAccountByName(getBaseQuery(), new_account.Name)
	if get_err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success":   true,
			"accountid": account[0].ID,
		})
		return
	}
	accountid, ins_err := insertAccount(new_account)
	if ins_err != nil {
		log.Print(ins_err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": ins_err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"success":   true,
		"accountid": accountid,
	})
}

func bindAccount(c *gin.Context) (account Account, err error) {
	var new_account Account
	if err := c.ShouldBind(&new_account); err != nil {
		return new_account, err
	}
	c.Bind(&new_account)
	return new_account, nil
}

func insertAccount(account Account) (id int64, err error) {
	var lastid int64
	query := "INSERT INTO account (account_name, account_type, rewards_features, payment_day, statement_day)\n"
	query += "VALUES (?, ?, ?, ?, ?)"

	res, err := DB.Exec(query,
		account.Name, account.Type, account.Features,
		account.PmtDay, account.StmtDay,
	)
	if err != nil {
		return lastid, err
	}
	lastid, err = res.LastInsertId()
	if err != nil {
		return lastid, err
	}
	return lastid, nil
}

func UpdateAccount(c *gin.Context) {

}

func DeleteAccount(c *gin.Context) {

}

func getBaseQuery() (query string) {
	base_query := "SELECT accountid, account_name, account_type, rewards_features, payment_day, statement_day\n"
	base_query += "FROM account\n"
	return base_query
}
