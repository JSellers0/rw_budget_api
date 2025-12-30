package services

type Transaction struct {
	ID        int    `json:"ID"`
	TransDate string `json:"trans_date"`
	CFDate    string `json:"cf_date"`
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
	return nil, nil
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
	//ToDo: buildGetTransQuery
	return "" + query_add
}
