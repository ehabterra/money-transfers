package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/ehabterra/money-transfers/internal/models"
)

func TestNewTransaction(t *testing.T) {
	type args struct {
		accountService *Account
		transactions   []models.Transaction
	}

	accountService, transactions := prepare(t)

	tests := []struct {
		name string
		args args
		want *Transaction
	}{
		{
			"basic",
			args{accountService, transactions},
			&Transaction{
				accountService: accountService,
				transactions:   transactions,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransaction(tt.args.accountService, tt.args.transactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func prepare(t *testing.T) (*Account, []models.Transaction) {
	account := NewAccount([]models.Account{
		{"1", 100, false},
		{"2", 200, false},
	})

	fromAccount, err := account.GetAccount("1")
	if err != nil {
		t.Fatalf("could not get from account: %v", err)
	}

	toAccount, err := account.GetAccount("2")
	if err != nil {
		t.Fatalf("could not get from account: %v", err)
	}

	transactions := []models.Transaction{
		{
			ID:          1,
			FromAccount: fromAccount,
			ToAccount:   toAccount,
			Amount:      50,
			Description: "test",
		},
	}
	return account, transactions
}

func TestTransaction_AddTransaction(t *testing.T) {
	type fields struct {
		accountService *Account
		transactions   []models.Transaction
	}
	type args struct {
		transaction TransactionData
	}

	accountService, transactions := prepare(t)
	fromAccount, err := accountService.GetAccount("1")
	if err != nil {
		t.Fatalf("error when getting from account: %v", err)
	}
	toAccount, err := accountService.GetAccount("2")
	if err != nil {
		t.Fatalf("error when getting from account: %v", err)
	}

	fromBalanceBefore := fromAccount.Balance
	toBalanceBefore := toAccount.Balance

	date := time.Now().UTC()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		{
			"basic",
			fields{
				accountService: accountService,
				transactions:   transactions,
			},
			args{
				TransactionData{
					FromAccount: "1",
					ToAccount:   "2",
					Amount:      50,
					Description: "test",
					Date:        date,
				},
			},
			&models.Transaction{
				ID:          2,
				FromAccount: fromAccount,
				ToAccount:   toAccount,
				Amount:      50,
				Description: "test",
				Date:        date,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Transaction{
				accountService: tt.fields.accountService,
				transactions:   tt.fields.transactions,
			}
			got, err := a.AddTransaction(tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTransaction() got = %v, want %v", got, tt.want)
			}
			if fromAccount.Balance != fromBalanceBefore-tt.want.Amount {
				t.Errorf("AddTransaction() from balance is not correct: got = %v, want %v", fromAccount.Balance, fromBalanceBefore-tt.want.Amount)
			}
			if toAccount.Balance != toBalanceBefore+tt.want.Amount {
				t.Errorf("AddTransaction() to balance is not correct: got = %v, want %v", toAccount.Balance, toBalanceBefore+tt.want.Amount)
			}
		})
	}
}

func TestTransaction_DeleteTransaction(t *testing.T) {
	type fields struct {
		accountService *Account
		transactions   []models.Transaction
	}
	type args struct {
		id int
	}
	accountService, transactions := prepare(t)

	fromBalanceBefore := transactions[0].FromAccount.Balance
	toBalanceBefore := transactions[0].ToAccount.Balance

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		{
			"basic",
			fields{
				accountService: accountService,
				transactions:   transactions,
			},
			args{id: 1},
			&transactions[0],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Transaction{
				accountService: tt.fields.accountService,
				transactions:   tt.fields.transactions,
			}
			got, err := a.DeleteTransaction(tt.args.id)
			if got == nil || (err != nil) != tt.wantErr {
				t.Errorf("DeleteTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && len(transactions) == 0 {
				t.Errorf("DeleteTransaction() got = %v, want %v", got, tt.want)
			}
			if got.FromAccount.Balance != fromBalanceBefore+tt.want.Amount {
				t.Errorf("AddTransaction() from balance is not correct: got = %v, want %v", got.FromAccount.Balance, fromBalanceBefore-tt.want.Amount)
			}
			if got.ToAccount.Balance != toBalanceBefore-tt.want.Amount {
				t.Errorf("AddTransaction() to balance is not correct: got = %v, want %v", got.ToAccount.Balance, toBalanceBefore+tt.want.Amount)
			}
		})
	}
}
func TestTransaction_GetTransaction(t *testing.T) {
	type fields struct {
		accountService *Account
		transactions   []models.Transaction
	}
	type args struct {
		id int
	}
	accountService, transactions := prepare(t)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		{
			"basic",
			fields{
				accountService: accountService,
				transactions:   transactions,
			},
			args{id: 1},
			&transactions[0],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Transaction{
				accountService: tt.fields.accountService,
				transactions:   tt.fields.transactions,
			}
			got, err := a.GetTransaction(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_GetTransactions(t *testing.T) {
	type fields struct {
		accountService *Account
		transactions   []models.Transaction
	}
	accountService, transactions := prepare(t)

	tests := []struct {
		name   string
		fields fields
		want   []models.Transaction
	}{
		{
			"basic",
			fields{
				accountService: accountService,
				transactions:   transactions,
			},
			transactions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Transaction{
				accountService: tt.fields.accountService,
				transactions:   tt.fields.transactions,
			}
			if got := a.GetTransactions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
