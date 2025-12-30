package handlers

import (
	s "rw_budget/api/services"

	"github.com/gin-gonic/gin"
)

type TransactionHanlder interface {
	GetTransactions(*gin.Context)
	GetTransactionByID(*gin.Context)
	PostTransaction(*gin.Context)
	PutTransaction(*gin.Context)
	DeleteTransaction(*gin.Context)
}

type transactionHandler struct {
	svc s.TransactionService
}

func NewTransactionHandler(service s.TransactionService) TransactionHanlder {
	return &transactionHandler{
		svc: service,
	}
}

func (h *transactionHandler) GetTransactions(c *gin.Context) {
	//ToDo: GetTransactions
}

func (h *transactionHandler) GetTransactionByID(c *gin.Context) {
	//ToDo: GetTransactionByID
}

func (h *transactionHandler) PostTransaction(c *gin.Context) {
	//ToDo: PostTransaction
}

func (h *transactionHandler) PutTransaction(c *gin.Context) {
	//ToDo: PutTransaction
}

func (h *transactionHandler) DeleteTransaction(c *gin.Context) {
	//ToDo: DeleteTransaction
}

func bindTransaction(c *gin.Context) (*s.Transaction, error) {
	//ToDo: bindTransaction
	return nil, nil
}
