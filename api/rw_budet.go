package main

import (
	"fmt"
	"log"
	"os"
	da "rw_budget/api/api/internal"

	"github.com/joho/godotenv"
)

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
}
