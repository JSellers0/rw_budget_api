package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	da "rw_budget/api/api/internal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type account struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var accounts = []account{
	{ID: 1, Name: "Test", Type: "Credit Card"},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	safe_pwd := os.Getenv("MSQLPWD")
	db := da.ConnectDB(safe_pwd)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/accounts", getAccounts)

	router.Run("localhost:8080")
}

func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}
