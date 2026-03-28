package services

import "database/sql"

type Services struct {
	Account     AccountService
	Cashflow    CashflowService
	Category    CategoryService
	Transaction TransactionService
}

func NewServices(db *sql.DB) *Services {
	return &Services{
		Account:     NewAccountService(db),
		Cashflow:    NewCashflowService(db),
		Category:    NewCategoryService(db),
		Transaction: NewTransactionService(db),
	}
}
