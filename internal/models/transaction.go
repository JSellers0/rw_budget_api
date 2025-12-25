package models

import "database/sql"

type Transaction struct {
	ID        int    `json:"ID"`
	TransDate string `json:"trans_date"`
	CFDate    string `json:"cf_date"`
}

func getTransactions(db *sql.DB) []Transaction {
	transactions := []Transaction{
		{ID: 1},
	}
	return transactions
}

func getTransaction(id int) Transaction {
	transaction := Transaction{ID: id}
	return transaction
}

func getTransactionsByDate(date string) []Transaction {
	transactions := []Transaction{
		{ID: 1},
	}
	return transactions
}