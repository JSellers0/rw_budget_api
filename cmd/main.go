package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"rw_budget_api/config"
	"rw_budget_api/database"
	"rw_budget_api/middleware"
	"rw_budget_api/routes"
	"rw_budget_api/services"
)

func main() {
	// Set up app configs
	log.Print("Loading config")
	config.Init()
	// Set up loggers
	accessLog, errorLog, err := middleware.SetupLogger(config.LogDir)
	if err != nil {
		log.Fatalf("Failed to setup logger: %v", err)
	}
	defer accessLog.(*os.File).Close()
	defer errorLog.(*os.File).Close()

	// Set up Database, Services, Routes
	if err := database.GetDB(); err != nil {
		log.Fatal(err)
	}

	svc := services.NewServices(database.DB)

	// Set Gin mode
	gin.SetMode(config.GinMode)
	svr := gin.New()
	svr.Use(gin.Recovery())
	svr.Use(middleware.AccessLogger(accessLog))
	svr.Use(middleware.ErrorLogger(errorLog))

	routes.SetupRoutesV1(svr, svc)

	apiAddress := config.ApiHost + ":" + config.ApiPort

	log.Printf("Starting server on %s", apiAddress)
	svr.Run(apiAddress)
}
