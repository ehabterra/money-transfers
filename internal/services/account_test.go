package services

import (
	"reflect"
	"testing"

	"github.com/ehabterra/money-transfers/internal/models"
)

func TestAccount_AddAccount(t *testing.T) {
	type fields struct {
		accounts []models.Account
	}
	type args struct {
		accountNo string
		balance   float32
	}

	account := &models.Account{AccountNo: "2", Balance: 15}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Account
		wantErr bool
	}{
		{"basic",
			fields{
				nil,
			},
			args{account.AccountNo, account.Balance},
			account,
			false,
		},
		{"empty_account",
			fields{
				[]models.Account{
					{"1", 0, false},
				},
			},
			args{"", account.Balance},
			nil,
			true,
		},
		{"wrong_account",
			fields{
				[]models.Account{
					{"1", 0, false},
				},
			},
			args{"1", account.Balance},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAccount(tt.fields.accounts)
			got, err := a.AddAccount(tt.args.accountNo, tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_DeleteAccount(t *testing.T) {
	type fields struct {
		accounts []models.Account
	}
	type args struct {
		accountNo string
	}
	account := &models.Account{AccountNo: "2", Balance: 15}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				[]models.Account{
					{"1", 0, false},
					*account,
				},
			},
			args:    args{account.AccountNo},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				accounts: tt.fields.accounts,
			}
			if err := a.DeleteAccount(tt.args.accountNo); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_GetAccount(t *testing.T) {
	type fields struct {
		accounts []models.Account
	}
	type args struct {
		accountNo string
	}
	account := &models.Account{AccountNo: "2", Balance: 15}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Account
		wantErr bool
	}{
		{"basic",
			fields{
				[]models.Account{
					{"1", 0, false},
					*account,
				},
			},
			args{account.AccountNo},
			account,
			false,
		},
		{"not_exists",
			fields{
				[]models.Account{
					{"1", 0, false},
					*account,
				},
			},
			args{"3"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAccount(tt.fields.accounts)
			got, err := a.GetAccount(tt.args.accountNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_GetAccounts(t *testing.T) {
	type fields struct {
		accounts []models.Account
	}
	accounts := []models.Account{
		{"1", 0, false},
		{"2", 15, false},
	}

	tests := []struct {
		name   string
		fields fields
		want   []models.Account
	}{
		{"basic",
			fields{
				accounts,
			},
			accounts,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				accounts: tt.fields.accounts,
			}
			if got := a.GetAccounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_UpdateAccount(t *testing.T) {
	type fields struct {
		accounts []models.Account
	}
	type args struct {
		accountNo string
		balance   float32
	}
	account := &models.Account{AccountNo: "2", Balance: 15}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Account
		wantErr bool
	}{
		{"basic",
			fields{
				[]models.Account{
					{"1", 0, false},
					{"2", 0, false},
				},
			},
			args{account.AccountNo, account.Balance},
			account,
			false,
		},
		{"not_exists",
			fields{
				[]models.Account{
					{"1", 0, false},
					{"2", 0, false},
				},
			},
			args{"3", account.Balance},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				accounts: tt.fields.accounts,
			}
			got, err := a.UpdateAccount(tt.args.accountNo, tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAccount(t *testing.T) {
	type args struct {
		accounts []models.Account
	}
	accounts := []models.Account{
		{"1", 0, false},
		{"2", 0, false},
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{"basic",
			args{
				accounts,
			},
			&Account{accounts},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
