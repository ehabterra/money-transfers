package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ehabterra/money_transfer/internal/multiplexer"

	"github.com/ehabterra/money_transfer/internal/services"
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	zeroTime := time.Time{}

	if transaction.Date == zeroTime {
		transaction.Date = time.Now()
	}

	transactionData, err := a.service.AddTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transactionData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	w.WriteHeader(http.StatusCreated)
}

func (a *Transaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(multiplexer.PARAMS)
	m := params.(map[string]*string)

	id, err := strconv.Atoi(*m["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	transactionData, err := a.service.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transactionData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (a *Transaction) GetTransactions(w http.ResponseWriter, _ *http.Request) {
	transactions := a.service.GetTransactions()

	data, err := json.Marshal(transactions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (a *Transaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(multiplexer.PARAMS)
	m := params.(map[string]*string)

	id, err := strconv.Atoi(*m["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = a.service.DeleteTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
