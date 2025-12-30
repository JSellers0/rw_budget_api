package services

type CashflowSummary struct {
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

type CashflowChart struct {
	Month         string
	ChartCategory string
	Amount        float64
}

type CashflowCardBalances struct{}

type CashflowService interface {
	ReadCashflowSummary(string, string) (*CashflowSummary, error)
	ReadCashflowChart(string, string, string) (*CashflowChart, error)
	ReadCashflowCardBalances(string, string) (*CashflowCardBalances, error)
}

type cashflowService struct{}

func NewCashflowService() CashflowService {
	return &cashflowService{}
}

func (s *cashflowService) ReadCashflowSummary(year string, month string) (*CashflowSummary, error) {
	var data CashflowSummary
	base_query := "SELECT cash_remain_sum, cash_in_sum, cash_out_sum\n"
	base_query += "\t, cash_remain_top, cash_in_top, cash_out_top\n"
	base_query += "\t, cash_remain_bot, cash_in_bot, cash_out_bot\n"
	base_query += "FROM vw_cashflow\nWHERE flow_year = ?\n"
	base_query += "\t AND flow_month = ?\n;"

	if err := DB.QueryRow(base_query, year, month).Scan(
		&data.SumRemain, &data.SumIn, &data.SumOut,
		&data.TopRemain, &data.TopIn, &data.TopOut,
		&data.BotRemain, &data.BotIn, &data.BotOut,
	); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *cashflowService) ReadCashflowChart(year string, month string, chart_limit string) (*CashflowChart, error) {
	month_start := year + "-" + month + "-01"
	chart_query := `
	WITH ccc AS (
		SELECT tran_month_name, cashflow_category, amount, tran_month_start
		FROM vw_cashflow_chart
		WHERE tran_month_start <= ?
		ORDER BY tran_month_start DESC
		LIMIT ` + chart_limit + `
	)
	SELECT tran_month_name, cashflow_category, amount
	FROM ccc
	ORDER BY tran_month_start ASC
	;
	`
	_, err := DB.Query(chart_query, month_start)
	if err != nil {
		return nil, err
	}
	//ToDo: Turn query data into CashflowChart
	return &CashflowChart{}, nil
}

func (s *cashflowService) ReadCashflowCardBalances(year string, month string) (*CashflowCardBalances, error) {
	//ToDo: ReadCashflowCardBalances
	return &CashflowCardBalances{}, nil
}
