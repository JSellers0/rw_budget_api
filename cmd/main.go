package main

import (
	h "rw_budget/api/handlers"
	s "rw_budget/api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	ah := h.NewAccountHandler(s.NewAccountService())
	ch := h.NewCategoryHandler(s.NewCategoryService())
	cfh := h.NewCashflowHandler(s.NewCashflowService())

	router := gin.Default()
	router.GET("/accounts", ah.GetAccounts)
	router.POST("/accounts", ah.PostAccount)
	router.GET("/accounts/:id", ah.GetAccountByID)
	router.PUT("/accounts/:id", ah.PutAccount)
	router.DELETE("/accounts/:id", ah.DeleteAccount)

	router.GET("/categories", ch.GetCategories)
	router.GET("/categories/:id", ch.GetCategoryByID)
	router.POST("/categories", ch.PostCategory)
	router.PUT("/categories", ch.PutCategory)
	router.DELETE("/categories/:id", ch.DeleteCategory)

	router.GET("cashflows/summary/:year/:month", cfh.GetCashflowSummary)
	router.GET("cashflows/chart/:year/:month", cfh.GetCashflowChart)
	router.GET("cashflows/card_balances/:year/:month", cfh.GetCashflowCardBalances)

	router.Run("localhost:8080")
}
