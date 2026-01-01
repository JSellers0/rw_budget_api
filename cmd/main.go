package main

import (
	"net/http"
	h "rw_budget/api/handlers"
	s "rw_budget/api/services"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ah := h.NewAccountHandler(s.NewAccountService())
	ch := h.NewCategoryHandler(s.NewCategoryService())
	cfh := h.NewCashflowHandler(s.NewCashflowService())
	th := h.NewTransactionHandler(s.NewTransactionService())

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "available",
			"timestamp": time.Now().Unix(),
		})
	})
	router.GET("/accounts", ah.GetAccounts)
	router.POST("/accounts", ah.PostAccount)
	router.GET("/accounts/:id", ah.GetAccountByID)
	router.PUT("/accounts/:id", ah.PutAccount)
	router.DELETE("/accounts/:id", ah.DeleteAccount)

	router.GET("/categories", ch.GetCategories)
	router.POST("/categories", ch.PostCategory)
	router.GET("/categories/:id", ch.GetCategoryByID)
	router.PUT("/categories/:id", ch.PutCategory)
	router.DELETE("/categories/:id", ch.DeleteCategory)

	router.GET("/transactions", th.GetTransactions)
	router.POST("/transactions", th.PostTransaction)
	router.GET("/transactions/:id", th.GetTransactionByID)
	router.PUT("/transactions/:id", th.PutTransaction)
	router.DELETE("/transactions/:id", th.DeleteTransaction)

	router.GET("cashflows/summary/:year/:month", cfh.GetCashflowSummary)
	router.GET("cashflows/chart/:year/:month", cfh.GetCashflowChart)
	router.GET("cashflows/card_balances/:year/:month", cfh.GetCashflowCardBalances)

	router.Run("localhost:8080")
}
