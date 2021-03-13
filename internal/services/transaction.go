package services

import (
	"errors"
	"time"

	"github.com/ehabterra/money-transfers/internal/models"
)

// TransactionData to transfer money between accounts
type TransactionData struct {
	ID          int       `json:"id"`
	FromAccount string    `json:"from_account"`
	ToAccount   string    `json:"to_account"`
	Amount      float32   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Deleted     bool
}

type Transaction struct {
	accountService *Account
	transactions   []models.Transaction
}

// NewTransaction construct transactions service
func NewTransaction(accountService *Account, transactions []models.Transaction) *Transaction {
	if transactions == nil {
		transactions = make([]models.Transaction, 0)
	}

	return &Transaction{accountService: accountService, transactions: transactions}
}

// AddTransaction creates new transaction
func (a *Transaction) AddTransaction(data TransactionData) (*models.Transaction, error) {
	if data.FromAccount == data.ToAccount {
		return nil, errors.New("cannot transfer to the same account")
	}

	fromAccount, err := a.accountService.GetAccount(data.FromAccount)
	if err != nil {
		return nil, err
	}

	toAccount, err := a.accountService.GetAccount(data.ToAccount)
	if err != nil {
		return nil, err
	}

	id := 1
	if len(a.transactions) > 0 {
		id = a.transactions[len(a.transactions)-1].ID + 1
	}

	trn, err := models.NewTransaction(id, fromAccount, toAccount, data.Amount, data.Description, data.Date)
	if err != nil {
		return nil, err
	}

	fromAccount.Balance -= data.Amount
	toAccount.Balance += data.Amount

	a.transactions = append(a.transactions, *trn)

	return trn, err
}

// DeleteTransaction deletes an transaction by transaction number
func (a *Transaction) DeleteTransaction(id int) (*models.Transaction, error) {
	transaction, err := a.GetTransaction(id)
	if err != nil {
		return nil, err
	}

	// return amounts to balance
	transaction.FromAccount.Balance += transaction.Amount
	transaction.ToAccount.Balance -= transaction.Amount

	transaction.Deleted = true
	return transaction, nil
}

// GetTransaction get an transaction data for specific transaction number
func (a *Transaction) GetTransaction(id int) (*models.Transaction, error) {

	for i := range a.transactions {
		trn := &a.transactions[i]
		if trn.ID == id && !trn.Deleted {
			return trn, nil
		}
	}

	return nil, errors.New("transaction number is not exists")
}

// GetTransactions get all transactions
func (a *Transaction) GetTransactions() []models.Transaction {
	transactions := make([]models.Transaction, 0)

	for _, trn := range a.transactions {
		if !trn.Deleted {
			transactions = append(transactions, trn)
		}
	}

	return transactions
}
