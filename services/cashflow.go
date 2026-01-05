package services

type SummaryData struct {
	FlowYear      int     `json:"flow_year"`
	FlowMonth     int     `json:"flow_month"`
	CashflowGroup string  `json:"cashflow_group"`
	MonthGroup    string  `json:"month_group"`
	Amount        float64 `json:"amount"`
}

type ChartData struct {
	Month         string  `json:"tran_month_name"`
	ChartCategory string  `json:"cashflow_category"`
	Amount        float64 `json:"amount"`
}

type CardBalance struct {
	FlowYear    int     `json:"flow_year"`
	FlowMonth   int     `json:"flow_month"`
	AccountID   int     `json:"accountid"`
	AccountName string  `json:"account_name"`
	Chg_bal     float64 `json:"chg_bal"`
	Pmt_bal     float64 `json:"pmt_bal"`
	Cur_bal     float64 `json:"cur_bal"`
	Pnd_bal     float64 `json:"pnd_bal"`
}

type CashflowService interface {
	ReadCashflowSummary(string, string) ([]*SummaryData, error)
	ReadCashflowChart(string, string, string) ([]*ChartData, error)
	ReadCashflowCardBalances(string, string) ([]*CardBalance, error)
}

type cashflowService struct{}

func NewCashflowService() CashflowService {
	return &cashflowService{}
}

func (s *cashflowService) ReadCashflowSummary(year string, month string) ([]*SummaryData, error) {
	var records []*SummaryData
	base_query := "SELECT flow_year, flow_month, cashflow_group, month_group, amount\n"
	base_query += "FROM vw_cashflow_summary_api\nWHERE flow_year = ?\n\t AND flow_month = ?\n;"

	res, err := DB.Query(base_query, year, month)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var data SummaryData
		if err := res.Scan(
			&data.FlowYear, &data.FlowMonth, &data.CashflowGroup,
			&data.MonthGroup, &data.Amount,
		); err != nil {
			return nil, err
		}
		records = append(records, &data)
	}
	return records, nil
}

func (s *cashflowService) ReadCashflowChart(year string, month string, chart_limit string) ([]*ChartData, error) {
	var records []*ChartData
	month_start := year + "-" + month + "-01"
	chart_query := `
	WITH ccc AS (
		SELECT tran_month_name, cashflow_category, amount, tran_month_start
		FROM vw_cashflow_chart
		WHERE tran_month_start <= ?
		ORDER BY tran_month_start DESC
		LIMIT ?
	)
	SELECT tran_month_name, cashflow_category, amount
	FROM ccc
	ORDER BY tran_month_start ASC
	;
	`
	res, err := DB.Query(chart_query, month_start, chart_limit)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var cfs ChartData
		if err := res.Scan(
			&cfs.Month, &cfs.ChartCategory, &cfs.Amount,
		); err != nil {
			return nil, err
		}
		records = append(records, &cfs)
	}
	return records, nil
}

func (s *cashflowService) ReadCashflowCardBalances(year string, month string) ([]*CardBalance, error) {
	var records []*CardBalance
	cb_query := "SELECT * FROM vw_cashflow_card_balances WHERE flow_year = ? AND flow_month = ?;"
	res, err := DB.Query(cb_query, year, month)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var data CardBalance
		if err := res.Scan(
			&data.FlowYear, &data.FlowMonth, &data.AccountID, &data.AccountName,
			&data.Chg_bal, &data.Pmt_bal, &data.Cur_bal, &data.Pnd_bal,
		); err != nil {
			return nil, err
		}
		records = append(records, &data)
	}
	return records, nil
}
