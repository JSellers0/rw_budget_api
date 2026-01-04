package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"accounts": records,
	})
}

func (h *accountHandler) GetAccountByID(c *gin.Context) {
	record, err := h.svc.ReadAccountByID(c.Param("id"))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Unable to locate the provided account.",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Unable to locate the provided account.",
			"error":   err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"accounts": record,
	})
}

func (h *accountHandler) PostAccount(c *gin.Context) {
	account, err := bindAccount(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error unpacking account.  Check your payload",
			"error":   err.Error(),
		})
		return
	}
	new_id, err := h.svc.CreateAccount(*account)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "Account already exists.",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error creating account.",
			"error":   err.Error(),
		})
		return
	}
	account.ID = strconv.FormatInt(*new_id, 10)
	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"message":  "Account created successfully",
		"accounts": account,
	})
}

func (h *accountHandler) PutAccount(c *gin.Context) {
	if c.Writer.Written() {
		fmt.Println("Headers written at start of handler!")
	}
	account, err := bindAccount(c)
	if c.Writer.Written() {
		fmt.Println("Headers written after bind.")
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error unpacking account.  Check your payload",
			"error":   err.Error(),
		})
		return
	}
	if err := h.svc.UpdateAccount(*account); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Unable to locate the provided account.",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Unable to update the provided account.",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Account Updated Successfully.",
	})
}

func (h *accountHandler) DeleteAccount(c *gin.Context) {
	if err := h.svc.DeleteAccount(c.Param("id")); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Unable to locate the provided account.",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Unable to locate the provided account.",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Account deleted successfully.",
	})
}

func bindAccount(c *gin.Context) (account *s.Account, err error) {
	var new_account s.Account
	if err = c.ShouldBind(&new_account); err != nil {
		return nil, err
	}
	if err = c.Bind(&new_account); err != nil {
		return nil, err
	}
	return &new_account, nil
}
