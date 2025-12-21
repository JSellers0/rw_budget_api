package models

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category struct {
	ID   int    `json:"categoryid"`
	Name string `json:"category_name"`
}

func GetCategories(c *gin.Context) {
	var status int
	var records []Category

	base_query := "SELECT categoryid, category_name\n"
	base_query += "FROM category\n"

	if c.Query("id") != "" {
		status, records = getCategoryByID(base_query, c.Query("id"))
	} else if c.Query("name") != "" {
		status, records = getCategoryByName(base_query, c.Query("name"))
	} else {
		status, records = getAllCategories(base_query)
	}
	c.IndentedJSON(status, records)
}

func getCategoryByID(base_query string, id string) (status int, record []Category) {
	base_query += "WHERE categoryid = ?\n;"
	log.Print(base_query)
	var category Category
	if err := DB.QueryRow(base_query, id).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		log.Print(err.Error())
		return handleSqlErr(err), []Category{}
	}
	return http.StatusOK, []Category{category}
}

func getCategoryByName(base_query string, name string) (status int, record []Category) {
	base_query += "WHERE category_name = ?\n;"
	var category Category
	if err := DB.QueryRow(base_query, name).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		return handleSqlErr(err), []Category{}
	}
	return http.StatusOK, []Category{category}
}

func getAllCategories(base_query string) (status int, record []Category) {
	categories := []Category{}
	results, err := DB.Query(base_query)
	if err != nil {
		return handleSqlErr(err), categories
	} else {
		for results.Next() {
			var category Category
			err = results.Scan(
				&category.ID, &category.Name,
			)
			if err != nil {
				return http.StatusInternalServerError, []Category{}
			}
			categories = append(categories, category)
		}
		return http.StatusOK, categories
	}
}
