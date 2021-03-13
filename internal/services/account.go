package services

import (
	"errors"

	"github.com/ehabterra/money-transfers/internal/models"
)

type Account struct {
	accounts []models.Account
}

// NewAccount construct accounts service
func NewAccount(accounts []models.Account) *Account {
	if accounts == nil {
		accounts = make([]models.Account, 0)
	}

	return &Account{accounts: accounts}
}

// AddAccount creates new account
func (a *Account) AddAccount(accountNo string, balance float32) (*models.Account, error) {
	if len(accountNo) == 0 {
		return nil, errors.New("account number is required")
	}

	acc, err := a.GetAccount(accountNo)
	if err == nil {
		return nil, errors.New("account already exists")
	}

	acc = &models.Account{
		AccountNo: accountNo,
		Balance:   balance,
	}
	a.accounts = append(a.accounts, *acc)

	return acc, nil
}

// UpdateAccount updates an account balance by account number
func (a *Account) UpdateAccount(accountNo string, balance float32) (*models.Account, error) {

	acc, err := a.GetAccount(accountNo)
	if err != nil {
		return nil, err
	}

	acc.Balance = balance
	return acc, nil
}

// DeleteAccount deletes an account by account number
func (a *Account) DeleteAccount(accountNo string) error {
	acc, err := a.GetAccount(accountNo)
	if err != nil {
		return err
	}

	acc.Deleted = true
	return nil
}

// GetAccount get an account data for specific account number
func (a *Account) GetAccount(accountNo string) (*models.Account, error) {

	for i := range a.accounts {
		acc := &a.accounts[i]
		if acc.AccountNo == accountNo && !acc.Deleted {
			return acc, nil
		}
	}

	return nil, errors.New("account number is not exists")
}

// GetAccounts get all accounts
func (a *Account) GetAccounts() []models.Account {
	accounts := make([]models.Account, 0)

	for _, acc := range a.accounts {
		if !acc.Deleted {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}
