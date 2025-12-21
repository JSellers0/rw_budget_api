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
	var records []Category
	var err error
	base_query := getBaseQuery()

	if c.Query("id") != "" {
		records, err = getCategoryByID(base_query, c.Query("id"))
	} else if c.Query("name") != "" {
		records, err = getCategoryByName(base_query, c.Query("name"))
	} else {
		records, err = getAllCategories(base_query)
	}
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(getErrStatus(err), records)
	}
	c.IndentedJSON(http.StatusOK, records)
}

func getCategoryByID(base_query string, id string) (record []Category, err error) {
	base_query += "WHERE categoryid = ?\n;"
	log.Print(base_query)
	var category Category
	if err := DB.QueryRow(base_query, id).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		return []Category{}, err
	}
	return []Category{category}, nil
}

func getCategoryByName(base_query string, name string) (record []Category, err error) {
	base_query += "WHERE category_name = ?\n;"
	var category Category
	if err := DB.QueryRow(base_query, name).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		return []Category{}, err
	}
	return []Category{category}, nil
}

func getAllCategories(base_query string) (record []Category, err error) {
	categories := []Category{}
	results, err := DB.Query(base_query)
	if err != nil {
		return categories, err
	} else {
		for results.Next() {
			var category Category
			err = results.Scan(
				&category.ID, &category.Name,
			)
			if err != nil {
				return categories, err
			}
			categories = append(categories, category)
		}
		return categories, err
	}
}

func CreateCategory(c *gin.Context) {

}

func UpdateCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}
