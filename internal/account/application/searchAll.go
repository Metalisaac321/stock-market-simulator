package application

import (
	"context"

	"github.com/Metalisaac321/stock-market-simulator/internal/account"
)

type SearchAllAccounts struct {
	repository account.AccountRepository
}

func NewSearchAllAccounts(repository account.AccountRepository) SearchAllAccounts {
	return SearchAllAccounts{
		repository: repository,
	}
}

func (useCase SearchAllAccounts) Execute(ctx context.Context) ([]account.Account, error) {
	accounts, _ := useCase.repository.SearchAll(ctx)

	return accounts, nil
}
