package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"rw_budget_api/middleware"
	"rw_budget_api/routes"
)

func main() {
	// Set up loggers
	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "./logs"
	}

	accessLog, errorLog, err := middleware.SetupLogger(logDir)
	if err != nil {
		log.Fatalf("Failed to setup logger: %v", err)
	}
	defer accessLog.(*os.File).Close()
	defer errorLog.(*os.File).Close()

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	svr := gin.New()
	svr.Use(gin.Recovery())
	svr.Use(middleware.AccessLogger(accessLog))
	svr.Use(middleware.ErrorLogger(errorLog))

	svr.GET("/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "available",
			"timestamp": time.Now().Unix(),
		})
	})
	routes.SetupRoutesV1(svr)
	app_port, is_set := os.LookupEnv("GIN_API_PORT")
	if !is_set {
		app_port = "8081"
	}
	host_name := "localhost:" + app_port

	log.Printf("Starting server on %s", host_name)
	svr.Run(host_name)
}
