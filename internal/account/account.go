package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// AccountId
var ErrInvalidAccountId = errors.New("Invalid AccountId")

type AccountId struct {
	value string
}

func NewAccountId(value string) (AccountId, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return AccountId{}, fmt.Errorf("%w: %s", ErrInvalidAccountId, value)
	}

	return AccountId{
		value: v.String(),
	}, nil
}
func (id AccountId) Value() string {
	return id.value
}

// AccountCash
var ErrInvalidAccountCash = errors.New("The account cash must be greater than 0")

type AccountCash struct {
	value uint
}

func NewAccountCash(value uint) (AccountCash, error) {
	if value <= 0 {
		return AccountCash{}, ErrInvalidAccountCash
	}

	return AccountCash{
		value: value,
	}, nil
}
func (cash AccountCash) Value() uint {
	return cash.value
}

// Account
type Account struct {
	id   AccountId
	cash AccountCash
}

func NewAccount(id string, cash uint) (Account, error) {
	idVO, err := NewAccountId(id)
	if err != nil {
		return Account{}, err
	}

	cashVO, err := NewAccountCash(cash)
	if err != nil {
		return Account{}, err
	}

	return Account{
		id:   idVO,
		cash: cashVO,
	}, nil
}

func (a Account) Id() AccountId {
	return a.id
}

func (a Account) Cash() AccountCash {
	return a.cash
}

// Repository
type AccountRepository interface {
	Save(ctx context.Context, account Account) error
	SearchAll(ctx context.Context) ([]Account, error)
}
