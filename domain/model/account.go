package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/satori/uuid.go"
	"time"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (a *Account) isValid() error {
	_, err := govalidator.ValidateStruct(a)

	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number, ownerName string) (*Account, error) {
	account := Account{
		Bank:      bank,
		Number:    number,
		OwnerName: ownerName,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &account, nil
}
