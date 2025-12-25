package models

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Account struct {
	ID       string `json:"accountid" form:"accountid"`
	Name     string `json:"account_name" form:"account_name" binding:"required"`
	Type     string `json:"account_type" form:"account_type" binding:"required"`
	Features string `json:"rewards_features" form:"rewards_features"`
	PmtDay   string `json:"payment_day" form:"payment_day"`
	StmtDay  string `json:"statement_day" form:"statement_day"`
}

func GetAccounts(c *gin.Context) {
	var records []Account
	var err error
	base_query := getBaseAccountQuery()

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
	account, get_err := getAccountByName(getBaseAccountQuery(), new_account.Name)
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

func bindAccount(c *gin.Context) (account *Account, err error) {
	var new_account Account
	if err := c.ShouldBind(&new_account); err != nil {
		return nil, err
	}
	c.Bind(&new_account)
	return &new_account, nil
}

func insertAccount(account *Account) (id *int64, err error) {
	var lastid int64
	query := "INSERT INTO account (account_name, account_type, rewards_features, payment_day, statement_day)\n"
	query += "VALUES (?, ?, ?, ?, ?)"

	res, err := DB.Exec(query,
		account.Name, account.Type, account.Features,
		account.PmtDay, account.StmtDay,
	)
	if err != nil {
		return nil, err
	}
	lastid, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &lastid, nil
}

func UpdateAccount(c *gin.Context) {
	mod_account, err := bindAccount(c)
	if err != nil || mod_account.ID == "" {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	accounts, get_err := getAccountByID(getBaseAccountQuery(), mod_account.ID)
	if get_err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": get_err.Error(),
		})
		return
	}
	mod_account = merge_account_changes(accounts[0], mod_account)
	query := "UPDATE ACCOUNT SET\n"
	query += "SET account_name=?, account_type=?, rewards_features=?, "
	query += "payment_day=?, statement_day=?\n"
	query += "WHERE accountid=?"

	_, up_err := DB.Exec(query,
		mod_account.Name, mod_account.Type, mod_account.Features,
		mod_account.PmtDay, mod_account.StmtDay, time.Now().UTC().Format(time.DateTime),
	)
	if up_err != nil {
		c.IndentedJSON(http.StatusInternalServerError, up_err.Error())
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "Record updated successfully.",
	})

}

func merge_account_changes(ex_account Account, mod_account *Account) (merge_account *Account) {
	if mod_account.Features == "" {
		mod_account.Features = ex_account.Features
	}
	if mod_account.PmtDay == "" {
		mod_account.PmtDay = ex_account.PmtDay
	}
	if mod_account.StmtDay == "" {
		mod_account.StmtDay = ex_account.StmtDay
	}
	return mod_account
}

func DeleteAccount(c *gin.Context) {
	accountid := c.Param("id")
	_, get_err := getAccountByID(getBaseAccountQuery(), accountid)
	if get_err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": get_err.Error(),
		})
		return
	}

	query := "DELETE FROM accounts WHERE accountid = ?;"
	_, del_err := DB.Exec(query, accountid)
	if del_err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": true,
			"message": del_err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Account successfully deleted.",
	})
}

func getBaseAccountQuery() (query string) {
	base_query := "SELECT accountid, account_name, account_type, rewards_features, payment_day, statement_day\n"
	base_query += "FROM account\n"
	return base_query
}