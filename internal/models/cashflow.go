package models

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Cashflow struct {
	SumRemain float64 `json:"cash_remain_sum"`
	SumIn     float64 `json:"cash_in_sum"`
	SumOut    float64 `json:"cash_out_sum"`
	TopRemain float64 `json:"cash_remain_top"`
	TopIn     float64 `json:"cash_in_top"`
	TopOut    float64 `json:"cash_out_top"`
	BotRemain float64 `json:"cash_remain_bot"`
	BotIn     float64 `json:"cash_in_bot"`
	BotOut    float64 `json:"cash_out_bot"`
}

func GetCashflow(c *gin.Context) {
	ctype := c.Param("type")
	year := c.Param("year")
	month := c.Param("month")

	if ctype == "summary" {
		getCashflowSummary(c, year, month)
		return
	} else if ctype == "chart" {
		getCashflowChart(c, year, month)
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": "Unsupported Cashflow type: " + ctype,
	})
}

func getCashflowSummary(c *gin.Context, year string, month string) {
	var cashflow Cashflow
	base_query := "SELECT cash_remain_sum, cash_in_sum, cash_out_sum\n"
	base_query += "\t, cash_remain_top, cash_in_top, cash_out_top\n"
	base_query += "\t, cash_remain_bot, cash_in_bot, cash_out_bot\n"
	base_query += "FROM vw_cashflow\nWHERE flow_year = ?\n"
	base_query += "\t AND flow_month = ?\n;"

	log.Print(base_query)

	if err := DB.QueryRow(base_query, year, month).Scan(
		&cashflow.SumRemain, &cashflow.SumIn, &cashflow.SumOut,
		&cashflow.TopRemain, &cashflow.TopIn, &cashflow.TopOut,
		&cashflow.BotRemain, &cashflow.BotIn, &cashflow.BotOut,
	); err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, cashflow)
}

func getCashflowChart(c *gin.Context, year string, month string) {
	limit_val, err := getChartLimit(c)
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	month_start := year + "-" + month + "-01"
	chart_query := `
	WITH ccc AS (
		SELECT tran_month_name, cashflow_category, amount, tran_month_start
		FROM vw_cashflow_chart
		WHERE tran_month_start <= ?
		ORDER BY tran_month_start DESC
		LIMIT ` + limit_val + `
	)
	SELECT tran_month_name, cashflow_category, amount
	FROM ccc
	ORDER BY tran_month_start ASC
	;
	`
	res, err := DB.Query(chart_query, month_start)
	if err != nil {
		log.Print(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
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