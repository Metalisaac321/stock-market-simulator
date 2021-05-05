package accounts

import (
	"net/http"

	"github.com/Metalisaac321/stock-market-simulator/internal/account/application"
	"github.com/gin-gonic/gin"
)

type accountDto struct {
	Id   string `json:"id"`
	Cash uint   `json:"cash"`
}
type searchAllAccountsResponse struct {
	Accounts []accountDto `json:"accounts"`
}

func SearchAllHandler(searchAllAccounts application.SearchAllAccounts) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accounts, err := searchAllAccounts.Execute(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		accountDtos := make([]accountDto, 1)

		for i, account := range accounts {
			accountDtos[i] = accountDto{
				Id:   account.Id().Value(),
				Cash: account.Cash().Value(),
			}
		}

		ctx.JSON(200, searchAllAccountsResponse{Accounts: accountDtos})
	}
}
