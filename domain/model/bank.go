package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/satori/uuid.go"
	"time"
)

type Bank struct {
	Base     `valid:"notnull"`
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (b *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(b)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}
