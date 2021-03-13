package models

// Account to manage accounts
type Account struct {
	AccountNo string  `json:"account_no"`
	Balance   float32 `json:"balance"`
	Deleted   bool    `json:"-"`
}

// NewAccount to construct new account
func NewAccount(accountNo string, balance float32) *Account {

	return &Account{
		AccountNo: accountNo,
		Balance:   balance,
	}
}
