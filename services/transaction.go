package services

type Transaction struct {
	TransactionID   int     `json:"transactionid"`
	TransactionDate string  `json:"transaction_date"`
	CashflowDate    string  `json:"cashflow_date"`
	MerchantName    string  `json:"merchant_name"`
	Amount          float32 `json:"amount"`
	CategoryID      string  `json:"categoryid"`
	CategoryName    string  `json:"category_name"`
	AccountID       int64   `json:"accountid"`
	AccountName     string  `json:"account_name"`
	TransactionType string  `json:"transaction_type"`
	Note            string  `json:"note"`
}

type TransactionService interface {
	CreateTransaction(Transaction) (*int64, error)
	ReadAllTransactions() ([]*Transaction, error)
	ReadTransactionByID(string) (*Transaction, error)
	ReadTransactionsByDate(string) ([]*Transaction, error)
	UpdateTransaction(string) error
	DeleteTransaction(string) error
}

type transactionService struct{}

func NewTransactionService() TransactionService {
	return &transactionService{}
}

func (s *transactionService) CreateTransaction(new_transaction Transaction) (*int64, error) {
	// ToDo: CreateTransaction
	return nil, nil
}

func (s *transactionService) ReadAllTransactions() ([]*Transaction, error) {
	// ToDo: ReadAllTransactions
	var records []*Transaction
	res, err := DB.Query(buildGetTransQuery(";"))
	if err != nil {
		return nil, err
	}
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

func (s *transactionService) ReadTransactionByID(id string) (*Transaction, error) {
	// ToDo: ReadTransactionByID
	return nil, nil
}

func (s *transactionService) ReadTransactionsByDate(filter_date string) ([]*Transaction, error) {
	// ToDo: ReadTransactionsByDate
	return nil, nil
}

func (s *transactionService) UpdateTransaction(string) error {
	// ToDo: UpdateTransaction
	return nil
}

func (s *transactionService) DeleteTransaction(string) error {
	// ToDo: DeleteTransaction
	return nil
}

func buildGetTransQuery(query_add string) string {
	query := "SELECT\n\tt.transactionid, t.transaction_date, t.cashflow_date, t.merchant_name, t.amount,\n"
	query += "\tt.transaction_type, t.note, t.accountid, t.categoryid, a.account_name, c.category_name\n"
	query += "FROM vw_transaction_detail\n"
	return query + query_add
}
