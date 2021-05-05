package application

import (
	"context"

	"github.com/Metalisaac321/stock-market-simulator/internal/account"
)

type CreateAccount struct {
	repository account.AccountRepository
}

func NewCreateAccount(repository account.AccountRepository) CreateAccount {
	return CreateAccount{
		repository: repository,
	}
}

func (useCase CreateAccount) Execute(ctx context.Context, id string, cash uint) error {
	course, err := account.NewAccount(id, cash)
	if err != nil {
		return err
	}
	return useCase.repository.Save(ctx, course)
}
