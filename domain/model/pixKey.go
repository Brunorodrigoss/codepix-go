package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// PixKeyRepositoryInterface ...
type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

// PixKey ...
type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

//  isValid ...
func (pixkey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixkey)

	if pixkey.Kind != "email" && pixkey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixkey.Status != "active" && pixkey.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

// NewPixKey ...
func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixkey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixkey.ID = uuid.NewV4().String()
	pixkey.CreatedAt = time.Now()

	err := pixkey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixkey, nil
}
