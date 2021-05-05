package accounts

import (
	"net/http"

	"github.com/Metalisaac321/stock-market-simulator/internal/account"
	"github.com/Metalisaac321/stock-market-simulator/internal/account/application"
	accountMap "github.com/Metalisaac321/stock-market-simulator/internal/account/mapper"
	"github.com/gin-gonic/gin"
)

type searchAllAccountsResponse struct {
	Accounts []accountMap.AccountDto `json:"accounts"`
}

func Map(accounts []account.Account, f func(account account.Account) accountMap.AccountDto) []accountMap.AccountDto {
	accountsDto := make([]accountMap.AccountDto, len(accounts))

	for i, account := range accounts {
		accountsDto[i] = f(account)
	}
	return accountsDto
}

func SearchAllHandler(searchAllAccounts application.SearchAllAccounts) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accounts, err := searchAllAccounts.Execute(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.JSON(200, searchAllAccountsResponse{Accounts: Map(accounts, accountMap.ToDto)})
	}
}
