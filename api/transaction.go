package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ehabterra/money-transfers/internal/multiplexer"

	"github.com/ehabterra/money-transfers/internal/services"
)

type Transaction struct {
	service *services.Transaction
}

func NewTransaction(service *services.Transaction) *Transaction {
	return &Transaction{service: service}
}

func (a *Transaction) AddTransaction(w http.ResponseWriter, r *http.Request) {
	transaction := services.TransactionData{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	zeroTime := time.Time{}

	if transaction.Date == zeroTime {
		transaction.Date = time.Now()
	}

	transactionData, err := a.service.AddTransaction(transaction)
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transactionData)
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
}

func (a *Transaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(multiplexer.PARAMS)
	m := params.(map[string]*string)

	id, err := strconv.Atoi(*m["id"])
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	transactionData, err := a.service.GetTransaction(id)
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transactionData)
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
}

func (a *Transaction) GetTransactions(w http.ResponseWriter, _ *http.Request) {
	transactions := a.service.GetTransactions()

	data, err := json.Marshal(transactions)
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
}

func (a *Transaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(multiplexer.PARAMS)
	m := params.(map[string]*string)

	id, err := strconv.Atoi(*m["id"])
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = a.service.DeleteTransaction(id)
	if err != nil {
		ShowError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
