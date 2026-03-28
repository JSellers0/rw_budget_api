package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	s "rw_budget_api/services"

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

func NewAccountHandler(accountService s.AccountService) AccountHandler {
	return &accountHandler{
		svc: accountService,
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
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Message: "Error retrieving accounts.",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    records,
	})
}

func (h *accountHandler) GetAccountByID(c *gin.Context) {
	record, err := h.svc.ReadAccountByID(c.Param("id"))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Message: "Unable to locate the provided account.",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Message: "Unable to locate the provided account.",
			Error:   err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    record,
	})
}

func (h *accountHandler) PostAccount(c *gin.Context) {
	var account s.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Message: "Error unpacking account.  Check your payload",
			Error:   err.Error(),
		})
		return
	}
	new_id, err := h.svc.CreateAccount(account)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusConflict, ErrorResponse{
				Success: false,
				Message: "Account already exists.",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Message: "Error creating account.",
			Error:   err.Error(),
		})
		return
	}
	account.ID = strconv.FormatInt(*new_id, 10)
	c.JSON(http.StatusCreated, SuccessResponse{
		Success: true,
		Message: "Account created successfully",
		Data:    account,
	})
}

func (h *accountHandler) PutAccount(c *gin.Context) {
	var account s.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Message: "Error unpacking account.  Check your payload",
			Error:   err.Error(),
		})
		return
	}
	if err := h.svc.UpdateAccount(account); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Message: "Unable to locate the provided account.",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Message: "Unable to update the provided account.",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: "Account Updated Successfully.",
	})
}

func (h *accountHandler) DeleteAccount(c *gin.Context) {
	if err := h.svc.DeleteAccount(c.Param("id")); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Message: "Unable to locate the provided account.",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Message: "Unable to locate the provided account.",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: "Account deleted successfully.",
	})
}
