package routes

import (
	h "rw_budget/api/handlers"
	s "rw_budget/api/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutesV1(r *gin.Engine) {
	setAccountRoutesV1(r)
	setCashflowRoutesV1(r)
	setCategoryRoutesV1(r)
	setTransactionRoutesV1(r)
}

func setAccountRoutesV1(g *gin.Engine) {
	ah := h.NewAccountHandler(s.NewAccountService())
	ar := g.Group("/v1/accounts")
	ar.GET("/", ah.GetAccounts)
	ar.POST("/", ah.PostAccount)
	ar.GET("/:id", ah.GetAccountByID)
	ar.PUT("/:id", ah.PutAccount)
	ar.DELETE("/:id", ah.DeleteAccount)
}

func setCashflowRoutesV1(g *gin.Engine) {
	cfh := h.NewCashflowHandler(s.NewCashflowService())
	cfr := g.Group("/v1/cashflows")
	cfr.GET("/summary/:year/:month", cfh.GetCashflowSummary)
	cfr.GET("/chart/:year/:month", cfh.GetCashflowChart)
	cfr.GET("/card_balances/:year/:month", cfh.GetCashflowCardBalances)
}

func setCategoryRoutesV1(g *gin.Engine) {
	ch := h.NewCategoryHandler(s.NewCategoryService())
	cr := g.Group("/v1/categories")
	cr.GET("/", ch.GetCategories)
	cr.POST("/", ch.PostCategory)
	cr.GET("/:id", ch.GetCategoryByID)
	cr.PUT("/:id", ch.PutCategory)
	cr.DELETE("/:id", ch.DeleteCategory)
}

func setTransactionRoutesV1(g *gin.Engine) {
	th := h.NewTransactionHandler(s.NewTransactionService())
	tr := g.Group("/v1/transactions")
	tr.GET("/", th.GetTransactions)
	tr.POST("/", th.PostTransaction)
	tr.GET("/:id", th.GetTransactionByID)
	tr.PUT("/:id", th.PutTransaction)
	tr.DELETE("/:id", th.DeleteTransaction)
}
