package models

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category struct {
	ID   int    `json:"categoryid" form:"categoryid"`
	Name string `json:"category_name" form:"category_name" binding:"required"`
}

func GetCategories(c *gin.Context) {
	var records []Category
	var err error
	base_query := getBaseCategoryQuery()

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

func getBaseCategoryQuery() string {
	return "SELECT categoryid, category_name\nFROM category\n"
}

func CreateCategory(c *gin.Context) {
	new_category, err := bindCategory(c)
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	categories, get_err := getCategoryByName(getBaseCategoryQuery(), new_category.Name)
	if get_err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success":   true,
			"accountid": categories[0].ID,
		})
		return
	}
	categoryid, ins_err := insertCategory(new_category)
	if ins_err != nil {
		log.Print(ins_err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": ins_err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"success":   true,
		"accountid": categoryid,
	})
}

func bindCategory(c *gin.Context) (category *Category, err error) {
	var new_category Category
	if err := c.ShouldBind(&new_category); err != nil {
		return nil, err
	}
	c.Bind(&new_category)
	return &new_category, nil
}

func insertCategory(new_category *Category) (id *int64, err error) {
	var lastid int64
	query := "INSERT INTO category (category_name) VALUES (?);"
	res, err := DB.Exec(query, new_category.Name)
	if err != nil {
		return nil, err
	}
	lastid, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &lastid, nil
}

func UpdateCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}