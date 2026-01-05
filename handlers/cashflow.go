package handlers

import (
	"log"
	"net/http"
	s "rw_budget/api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CashflowHandler interface {
	GetCashflowSummary(*gin.Context)
	GetCashflowChart(*gin.Context)
	GetCashflowCardBalances(*gin.Context)
}

type cashflowHandler struct {
	svc s.CashflowService
}

func NewCashflowHandler(cf_service s.CashflowService) CashflowHandler {
	return &cashflowHandler{
		svc: cf_service,
	}
}

func (h *cashflowHandler) GetCashflowSummary(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")

	data, err := h.svc.ReadCashflowSummary(year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"cashflows": data,
	})
}

func (h *cashflowHandler) GetCashflowChart(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	limit_val, err := getChartLimit(c)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	data, err := h.svc.ReadCashflowChart(year, month, limit_val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"chart":   data,
	})

}
func (h *cashflowHandler) GetCashflowCardBalances(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	data, err := h.svc.ReadCashflowCardBalances(year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

func getChartLimit(c *gin.Context) (limit string, err error) {
	var month_range int
	var chart_segs int

	if c.Query("month_range") == "" {
		month_range = 6
	} else {
		i, err := strconv.Atoi(c.Query("month_range"))
		if err != nil {
			return "", err
		}
		month_range = i
	}

	if c.Query("chart_segments") == "" {
		chart_segs = 3
	} else {
		i, err := strconv.Atoi(c.Query("chart_segments"))
		if err != nil {
			return "", err
		}
		chart_segs = i
	}

	return strconv.Itoa(month_range * chart_segs), nil

}
