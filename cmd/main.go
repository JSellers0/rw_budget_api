package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	router "rw_budget_api/routes"
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
	app_port, is_set := os.LookupEnv("GIN_API_PORT")
	if !is_set {
		app_port = "8081"
	}
	host_name := "localhost:" + app_port

	svr.Run(host_name)
}
