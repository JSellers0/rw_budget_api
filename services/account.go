package services

import (
	"time"
)

type Account struct {
	ID       string `json:"accountid" form:"accountid"`
	Name     string `json:"account_name" form:"account_name" binding:"required"`
	Type     string `json:"account_type" form:"account_type" binding:"required"`
	Features string `json:"rewards_features" form:"rewards_features"`
	PmtDay   string `json:"payment_day" form:"payment_day"`
	StmtDay  string `json:"statement_day" form:"statement_day"`
}

type AccountService interface {
	CreateAccount(Account) (*int64, error)
	ReadAccountByID(string) (*Account, error)
	ReadAccountsByName(string) ([]*Account, error)
	ReadAllAccounts() ([]*Account, error)
	UpdateAccount(Account) error
	DeleteAccount(string) error
}

type accountService struct{}

func NewAccountService() AccountService {
	return &accountService{}
}

func (s *accountService) CreateAccount(account Account) (id *int64, err error) {
	var lastid int64
	query := "INSERT INTO account (account_name, account_type, rewards_features, payment_day, statement_day)\n"
	query += "VALUES (?, ?, ?, ?, ?)"

	res, err := DB.Exec(query,
		account.Name, account.Type, account.Features,
		account.PmtDay, account.StmtDay,
	)
	if err != nil {
		return nil, err
	}
	lastid, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &lastid, nil
}

func (s *accountService) ReadAccountByID(id string) (*Account, error) {
	query := buildGetAccountQuery("WHERE accountid = ?\n;")
	var account Account
	if err := DB.QueryRow(
		query, id).Scan(
		&account.ID,
		&account.Name,
		&account.Type,
		&account.Features,
		&account.PmtDay,
		&account.StmtDay,
	); err != nil {
		return nil, err
	}
	return &account, nil
}

func (s *accountService) ReadAccountsByName(name string) (records []*Account, err error) {
	query := buildGetAccountQuery("WHERE account_name = ?\n;")
	accounts := []*Account{}
	results, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var account Account
		err = results.Scan(
			&account.ID,
			&account.Name,
			&account.Type,
			&account.Features,
			&account.PmtDay,
			&account.StmtDay,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (s *accountService) ReadAllAccounts() (records []*Account, err error) {
	accounts := []*Account{}
	results, err := DB.Query(buildGetAccountQuery(";"))
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var account Account
		err = results.Scan(
			&account.ID,
			&account.Name,
			&account.Type,
			&account.Features,
			&account.PmtDay,
			&account.StmtDay,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (s *accountService) UpdateAccount(mod_account Account) error {
	account, get_err := s.ReadAccountByID(mod_account.ID)
	if get_err != nil {
		return get_err
	}
	mod_account = merge_account_changes(account, mod_account)
	query := "UPDATE ACCOUNT SET\n"
	query += "SET account_name=?, account_type=?, rewards_features=?, "
	query += "payment_day=?, statement_day=?\n"
	query += "WHERE accountid=?"

	_, up_err := DB.Exec(query,
		mod_account.Name, mod_account.Type, mod_account.Features,
		mod_account.PmtDay, mod_account.StmtDay, time.Now().UTC().Format(time.DateTime),
	)
	if up_err != nil {
		return up_err
	}
	return nil
}

func (s *accountService) DeleteAccount(accountid string) error {
	_, get_err := s.ReadAccountByID(accountid)
	if get_err != nil {
		return get_err

	}
	query := "DELETE FROM accounts WHERE accountid = ?;"
	_, del_err := DB.Exec(query, accountid)
	if del_err != nil {
		return del_err
	}

	return nil
}

func buildGetAccountQuery(query_add string) (query string) {
	base_query := "SELECT accountid, account_name, account_type, rewards_features, payment_day, statement_day\n"
	base_query += "FROM account\n"
	return base_query + query_add
}

func merge_account_changes(ex_account *Account, mod_account Account) (merge_account Account) {
	if mod_account.Features == "" {
		mod_account.Features = ex_account.Features
	}
	if mod_account.PmtDay == "" {
		mod_account.PmtDay = ex_account.PmtDay
	}
	if mod_account.StmtDay == "" {
		mod_account.StmtDay = ex_account.StmtDay
	}
	return mod_account
}
