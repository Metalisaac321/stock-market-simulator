package creating

import (
	"context"

	mooc "github.com/Metalisaac321/stock-market-simulator/internal"
)

type AccountService struct {
	accountRepository mooc.AccountRepository
}

func NewAccountService(accountRepository mooc.AccountRepository) AccountService {
	return AccountService{
		accountRepository: accountRepository,
	}
}

func (service AccountService) CreateAccount(ctx context.Context, id string, cash uint) error {
	course, err := mooc.NewAccount(id, cash)
	if err != nil {
		return err
	}
	return service.accountRepository.Save(ctx, course)
}
