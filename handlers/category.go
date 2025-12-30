package handlers

import (
	"log"
	"net/http"
	s "rw_budget/api/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	GetCategories(*gin.Context)
	GetCategoryByID(*gin.Context)
	PostCategory(*gin.Context)
	PutCategory(*gin.Context)
	DeleteCategory(*gin.Context)
}

type categoryHandler struct {
	svc s.CategoryService
}

func NewCategoryHandler(category_service s.CategoryService) CategoryHandler {
	return &categoryHandler{
		svc: category_service,
	}
}

func (h categoryHandler) GetCategories(c *gin.Context) {
	var records []*s.Category
	var err error

	if c.Query("name") != "" {
		records, err = h.svc.ReadCategoriesByName(c.Query("name"))
	} else {
		records, err = h.svc.ReadAllCategories()
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Unable to locate records based on your request.",
		})
	}
	c.JSON(http.StatusOK, records)
}

func (h categoryHandler) GetCategoryByID(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Request did not include an id in the path.",
		})
	}

	record, err := h.svc.ReadCategoryByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Request did not include an id in the path.",
		})
	}
	c.JSON(http.StatusOK, record)
}

func (h categoryHandler) PostCategory(c *gin.Context) {
	new_category, err := bindCategory(c)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	categories, get_err := h.svc.ReadCategoriesByName(new_category.Name)
	if get_err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":   true,
			"accountid": categories[0].ID,
		})
		return
	}
	new_id, err := h.svc.CreateCategory(*new_category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, new_id)
}

func (h categoryHandler) PutCategory(c *gin.Context) {
	category, err := bindCategory(c)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "",
			"error":   err.Error(),
		})
		return
	}
	if err = h.svc.UpdateCategory(*category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category Updated Successfully.",
	})

}

func (h categoryHandler) DeleteCategory(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Request did not include an id in the path.",
		})
	}

	if err := h.svc.DeleteCategory(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error encountered deleting Category.",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category deleted successfully.",
	})

}

func bindCategory(c *gin.Context) (category *s.Category, err error) {
	var new_category s.Category
	if err := c.ShouldBind(&new_category); err != nil {
		return nil, err
	}
	c.Bind(&new_category)
	return &new_category, nil
}
