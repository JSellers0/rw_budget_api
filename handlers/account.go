package handlers

import (
	"log"
	"net/http"
	"strconv"

	s "rw_budget/api/services"

	"github.com/gin-gonic/gin"
)

type AccountHandler interface {
	GetAccounts(*gin.Context)
	GetAccountByID(*gin.Context)
	PostAccount(*gin.Context)
	PutAccount(*gin.Context)
	DeleteAccount(*gin.Context)
}

type accountHandler struct {
	svc s.AccountService
}

func NewAccountHandler(account_service s.AccountService) AccountHandler {
	return &accountHandler{
		svc: account_service,
	}
}

func bindAccount(c *gin.Context) (account *s.Account, err error) {
	var new_account s.Account
	if err := c.ShouldBind(&new_account); err != nil {
		return nil, err
	}
	c.Bind(&new_account)
	return &new_account, nil
}

func (h *accountHandler) GetAccounts(c *gin.Context) {
	var records []*s.Account
	var err error

	if c.Query("name") != "" {
		records, err = h.svc.ReadAccountsByName(c.Query("name"))
	} else {
		records, err = h.svc.ReadAllAccounts()
	}

	if err != nil {
		log.Print("Error retrieving accounts.")
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error retrieving accounts.",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, records)
}

func (h *accountHandler) GetAccountByID(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "URL Path is missing ID parameter",
		})
		return
	}
	record, err := h.svc.ReadAccountByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Unable to locate the provided account.",
		})
		return
	}
	c.JSON(http.StatusOK, record)
}

func (h *accountHandler) PostAccount(c *gin.Context) {
	account, err := bindAccount(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	new_id, err := h.svc.CreateAccount(*account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Account created successfully",
		"data":    `{"accountid": ` + strconv.Itoa(int(*new_id)) + `}`,
	})
}

func (h *accountHandler) PutAccount(c *gin.Context) {
	account, err := bindAccount(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	if err = h.svc.UpdateAccount(*account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Account Updated Successfully.",
	})
}

func (h *accountHandler) DeleteAccount(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "URL Path is missing ID parameter",
		})
		return
	}
	if err := h.svc.DeleteAccount(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Unable to locate the provided account.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Account deleted successfully.",
	})
}
