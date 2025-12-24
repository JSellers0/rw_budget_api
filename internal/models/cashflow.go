package models

import (
	"log"
	"net/http"

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
	var cashflow Cashflow
	year := c.Param("year")
	month := c.Param("month")

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
