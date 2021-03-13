package models

import (
	"errors"
	"fmt"
	"time"
)

// Transaction to transfer money between accounts
type Transaction struct {
	ID          int       `json:"id"`
	FromAccount *Account  `json:"from_account"`
	ToAccount   *Account  `json:"to_account"`
	Amount      float32   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Deleted     bool      `json:"-"`
}

// NewTransaction to construct new transaction
func NewTransaction(id int, from *Account, to *Account, amount float32, description string, date time.Time) (*Transaction, error) {

	if from == nil {
		return nil, errors.New("from account is not correct")
	}

	if to == nil {
		return nil, errors.New("to account is not correct")
	}

	if amount <= 0 {
		return nil, errors.New("amount should be greater than 0")
	}

	if amount > from.Balance {
		return nil, errors.New(fmt.Sprintf("no sufficient balance to proceed the transaction: got = %v, want = %v", from.Balance, amount))
	}

	return &Transaction{
		ID:          id,
		FromAccount: from,
		ToAccount:   to,
		Amount:      amount,
		Description: description,
		Date:        date.UTC(),
	}, nil
}
