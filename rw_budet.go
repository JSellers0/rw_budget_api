package main

import (
	"fmt"
	"log"
	da "rw_budget/api/internal"
	models "rw_budget/api/models"

	"github.com/gin-gonic/gin"
)

var DB = da.GetDB()

func main() {
	pingErr := da.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/accounts", models.GetAccounts)

	router.Run("localhost:8080")
}
