package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	router "rw_budget/api/routes"
)

func main() {
	svr := gin.Default()
	svr.GET("/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "available",
			"timestamp": time.Now().Unix(),
		})
	})
	router.SetupRoutesV1(svr)

	svr.Run("localhost:8080")
}
