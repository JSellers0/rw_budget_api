package main

import (
	"fmt"
	"log"
	db "rw_budget/api/internal/database"
	models "rw_budget/api/internal/models"

	"github.com/gin-gonic/gin"
)

var DB = db.GetDB()

func main() {
	pingErr := db.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/accounts", models.GetAccounts)
	router.POST("/accounts", models.CreateAccount)
	router.PUT("/accounts", models.UpdateAccount)
	router.DELETE("/accounts", models.DeleteAccount)

	router.GET("/categories", models.GetCategories)
	router.POST("/categories", models.CreateCategory)
	router.PUT("/categories", models.UpdateCategory)
	router.DELETE("/categories", models.DeleteCategory)

	router.GET("cashflow/:year/:month", models.GetCashflow)

	router.Run("localhost:8080")
}
