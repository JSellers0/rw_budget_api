package models

type Account struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Features string `json:"features"`
	PmtDate  int    `json:"pmt_date"`
	StmtDate int    `json:"stmt_date"`
}

func getAccounts() []Account {
	accounts := []Account{
		{ID: 1, Name: "Test"},
	}
	return accounts
}

func getAccountID(name string) int {
	account := Account{ID: 1, Name: name}
	return account.ID
}

func getAccount(id int) Account {
	account := Account{ID: id, Name: "test"}
	return account
}

func createAccount(deets Account) int {
	return 0
}

func updateAccount(deets Account) int {
	return 0
}

func deleteAccount(id int) int {
	isAuth := true
	if isAuth {
		return 0
	}
	return 1
}
