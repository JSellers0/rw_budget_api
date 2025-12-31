package services

import (
	"database/sql"
	"strconv"
)

type Transaction struct {
	TransactionID   int     `json:"transactionid"`
	TransactionDate string  `json:"transaction_date" binding:"required"`
	CashflowDate    string  `json:"cashflow_date" binding:"required"`
	MerchantName    string  `json:"merchant_name" binding:"required"`
	Amount          float32 `json:"amount" binding:"required"`
	CategoryID      string  `json:"categoryid" binding:"required"`
	CategoryName    string  `json:"category_name"`
	AccountID       int64   `json:"accountid" binding:"required"`
	AccountName     string  `json:"account_name"`
	TransactionType string  `json:"transaction_type" binding:"required"`
	Note            string  `json:"note"`
}

type TransactionService interface {
	CreateTransaction(Transaction) (*int64, error)
	ReadAllTransactions() ([]*Transaction, error)
	ReadTransactionByID(string) (*Transaction, error)
	ReadTransactionsByDateRange(string, string) ([]*Transaction, error)
	UpdateTransaction(Transaction) error
	DeleteTransaction(string) error
}

type transactionService struct{}

func NewTransactionService() TransactionService {
	return &transactionService{}
}

func (s *transactionService) CreateTransaction(new_trans Transaction) (*int64, error) {
	query := "INSERT INTO transactions (transaction_date, cashflow_date, transaction_type, merchant_name, amount, accountid, categoryid, note)\n"
	query += "VALUES (?,?,?,?,?,?,?,?)"
	res, err := DB.Exec(query, new_trans.TransactionDate, new_trans.CashflowDate, new_trans.TransactionType, new_trans.MerchantName,
		new_trans.Amount, new_trans.AccountID, new_trans.CategoryID, new_trans.Note,
	)
	if err != nil {
		return nil, err
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &lastid, nil
}

func (s *transactionService) ReadAllTransactions() ([]*Transaction, error) {
	res, err := DB.Query(buildGetTransQuery(";"))
	if err != nil {
		return nil, err
	}
	records, err := packageRows(res)
	if err := res.Scan(); err != nil {
		return nil, err
	}
	return records, nil
}

func (s *transactionService) ReadTransactionByID(id string) (*Transaction, error) {
	var data Transaction
	query := buildGetTransQuery("WHERE transactionid = ?")
	if err := DB.QueryRow(query, id).Scan(
		&data.TransactionID, &data.TransactionDate, &data.CashflowDate, &data.MerchantName,
		&data.Amount, &data.TransactionType, &data.Note, &data.AccountID, &data.AccountName,
		&data.CategoryID, &data.CategoryName,
	); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *transactionService) ReadTransactionsByDateRange(start_date string, end_date string) ([]*Transaction, error) {
	query := buildGetTransQuery("WHERE cashflow_date BETWEEN ? AND ?;")
	res, err := DB.Query(query, start_date, end_date)
	if err != nil {
		return nil, err
	}
	records, err := packageRows(res)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *transactionService) UpdateTransaction(mod_trans Transaction) error {
	_, get_err := s.ReadTransactionByID(strconv.Itoa(mod_trans.TransactionID))
	if get_err != nil {
		return get_err
	}
	// ToDo: Set update date and update by
	query := "UPDATE transactions SET\n"
	query += "transaction_date=?, cashflow_date=?, transaction_type=?, merchant_name=?,\n"
	query += "amount=?, accountid=?, categoryid=?, note=?\n"
	query += "WHERE transactionid = ?;"
	_, err := DB.Exec(query, mod_trans.TransactionDate, mod_trans.CashflowDate, mod_trans.TransactionType,
		mod_trans.MerchantName, mod_trans.Amount, mod_trans.AccountID, mod_trans.CategoryID, mod_trans.Note,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *transactionService) DeleteTransaction(id string) error {
	_, get_err := s.ReadTransactionByID(id)
	if get_err != nil {
		return get_err

	}
	query := "DELETE FROM transactions WHERE transactionid = ?;"
	_, del_err := DB.Exec(query, id)
	if del_err != nil {
		return del_err
	}
	return nil
}

func buildGetTransQuery(query_add string) string {
	query := "SELECT\n\ttransactionid, transaction_date, cashflow_date, merchant_name, amount,\n"
	query += "\ttransaction_type, note, accountid, categoryid, account_name, category_name\n"
	query += "FROM vw_transaction_detail\n"
	return query + query_add
}

func packageRows(res *sql.Rows) ([]*Transaction, error) {
	var records []*Transaction
	for res.Next() {
		var data Transaction
		if err := res.Scan(
			&data.TransactionID, &data.TransactionDate, &data.CashflowDate, &data.MerchantName,
			&data.Amount, &data.TransactionType, &data.Note, &data.AccountID, &data.AccountName,
			&data.CategoryID, &data.CategoryName,
		); err != nil {
			return nil, err
		}
		records = append(records, &data)
	}
	return records, nil
}
