package handlers

import (
	"net/http"
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
	var records []*s.Transaction
	var err error
	if c.Query("start_date") != "" && c.Query("end_date") != "" {
		records, err = h.svc.ReadTransactionsByDateRange(c.Query("start_date"), c.Query("end_date"))
	} else {
		records, err = h.svc.ReadAllTransactions()
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error returning records for your request.",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    records,
	})
}

func (h *transactionHandler) GetTransactionByID(c *gin.Context) {
	records, err := h.svc.ReadTransactionByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error returning records for your request.",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    records,
	})
}

func (h *transactionHandler) PostTransaction(c *gin.Context) {
	new_t, err := bindTransaction(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error unpacking transaction. Check your payload.",
			"error":   err.Error(),
		})
	}
	new_id, err := h.svc.CreateTransaction(*new_t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error inserting transaction.  Check your payload.",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"success":       true,
		"message":       "Transaction created successfully",
		"transactionid": new_id,
	})
}

func (h *transactionHandler) PutTransaction(c *gin.Context) {
	mod_t, err := bindTransaction(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error unpacking transaction. Check your payload.",
			"error":   err.Error(),
		})
	}
	if err := h.svc.UpdateTransaction(*mod_t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error updating transaction.  Check your payload.",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "TrTransaction Updated Successfully.",
	})

}

func (h *transactionHandler) DeleteTransaction(c *gin.Context) {
	if err := h.svc.DeleteTransaction(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error returning records for your request.",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func bindTransaction(c *gin.Context) (*s.Transaction, error) {
	var new_trans s.Transaction
	if err := c.ShouldBind(&new_trans); err != nil {
		return nil, err
	}
	c.Bind(&new_trans)
	return &new_trans, nil
}
