package api

import (
	"encoding/json"
	"net/http"

	"github.com/ehabterra/money-transfers/internal/multiplexer"

	"github.com/ehabterra/money-transfers/internal/models"
	"github.com/ehabterra/money-transfers/internal/services"
)

type Account struct {
	service *services.Account
}

func NewAccount(service *services.Account) *Account {
	return &Account{service: service}
}

func (a *Account) AddAccount(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accountData, err := a.service.AddAccount(account.AccountNo, account.Balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(accountData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
}

func (a *Account) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(multiplexer.PARAMS)
	m := params.(map[string]*string)

	account := models.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accountData, err := a.service.UpdateAccount(*m["accountNo"], account.Balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(accountData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
}

func (a *Account) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(multiplexer.PARAMS)
	m := params.(map[string]*string)

	accountData, err := a.service.GetAccount(*m["accountNo"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(accountData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
}

func (a *Account) GetAccounts(w http.ResponseWriter, _ *http.Request) {
	accounts := a.service.GetAccounts()

	data, err := json.Marshal(accounts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
}

func (a *Account) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(multiplexer.PARAMS)
	m := params.(map[string]*string)

	err := a.service.DeleteAccount(*m["accountNo"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
