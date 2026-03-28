package routes

import (
	"net/http"
	h "rw_budget_api/handlers"
	"rw_budget_api/services"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutesV1(r *gin.Engine, s *services.Services) {
	setHealthRouteV1(r)
	setAccountRoutesV1(r, s)
	setCashflowRoutesV1(r, s)
	setCategoryRoutesV1(r, s)
	setTransactionRoutesV1(r, s)
}

func setHealthRouteV1(g *gin.Engine) {
	g.GET("/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "available",
			"timestamp": time.Now().Unix(),
		})
	})
}

func setAccountRoutesV1(g *gin.Engine, s *services.Services) {
	ah := h.NewAccountHandler(s.Account)
	ar := g.Group("/v1/accounts")
	ar.GET("/", ah.GetAccounts)
	ar.POST("/", ah.PostAccount)
	ar.GET("/:id", ah.GetAccountByID)
	ar.PUT("/:id", ah.PutAccount)
	ar.DELETE("/:id", ah.DeleteAccount)
}

func setCashflowRoutesV1(g *gin.Engine, s *services.Services) {
	cfh := h.NewCashflowHandler(s.Cashflow)
	cfr := g.Group("/v1/cashflows")
	cfr.GET("/summary/:year/:month", cfh.GetCashflowSummary)
	cfr.GET("/chart/:year/:month", cfh.GetCashflowChart)
	cfr.GET("/card_balances/:year/:month", cfh.GetCashflowCardBalances)
}

func setCategoryRoutesV1(g *gin.Engine, s *services.Services) {
	ch := h.NewCategoryHandler(s.Category)
	cr := g.Group("/v1/categories")
	cr.GET("/", ch.GetCategories)
	cr.POST("/", ch.PostCategory)
	cr.GET("/:id", ch.GetCategoryByID)
	cr.PUT("/:id", ch.PutCategory)
	cr.DELETE("/:id", ch.DeleteCategory)
}

func setTransactionRoutesV1(g *gin.Engine, s *services.Services) {
	th := h.NewTransactionHandler(s.Transaction)
	tr := g.Group("/v1/transactions")
	tr.GET("/", th.GetTransactions)
	tr.POST("/", th.PostTransaction)
	tr.GET("/:id", th.GetTransactionByID)
	tr.PUT("/:id", th.PutTransaction)
	tr.DELETE("/:id", th.DeleteTransaction)
}
