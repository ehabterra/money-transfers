package main

import (
	"net/http"

	"github.com/ehabterra/money-transfers/api"
	"github.com/ehabterra/money-transfers/internal/multiplexer"
	"github.com/ehabterra/money-transfers/internal/services"
)

func prepareRoutes() *multiplexer.Server {
	// get instance from multiplexer server
	s := multiplexer.NewServer()

	// new account service
	accountService := services.NewAccount(nil)

	// new account API
	accountAPI := api.NewAccount(accountService)

	// account routes
	s.AddRoute("/accounts", http.MethodGet, accountAPI.GetAccounts)
	s.AddRoute("/accounts", http.MethodPost, accountAPI.AddAccount)
	s.AddRoute("/accounts/{accountNo}", http.MethodPut, accountAPI.UpdateAccount)
	s.AddRoute("/accounts/{accountNo}", http.MethodDelete, accountAPI.DeleteAccount)
	s.AddRoute("/accounts/{accountNo}", http.MethodGet, accountAPI.GetAccount)

	// new transaction service
	transactionService := services.NewTransaction(accountService, nil)

	// new transaction API
	transactionAPI := api.NewTransaction(transactionService)

	// transaction routes
	s.AddRoute("/transactions/", http.MethodGet, transactionAPI.GetTransactions)
	s.AddRoute("/transactions", http.MethodPost, transactionAPI.AddTransaction)
	s.AddRoute("/transactions/{id}", http.MethodDelete, transactionAPI.DeleteTransaction)
	s.AddRoute("/transactions/{id}", http.MethodGet, transactionAPI.GetTransaction)

	return s
}
