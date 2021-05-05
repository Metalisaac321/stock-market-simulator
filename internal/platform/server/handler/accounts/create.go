package accounts

import (
	"errors"
	"net/http"

	"github.com/Metalisaac321/stock-market-simulator/internal/account"
	accountApplication "github.com/Metalisaac321/stock-market-simulator/internal/account/application"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Id   string `json:"id"`
	Cash uint   `json:"cash"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(createAccount accountApplication.CreateAccount) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := createAccount.Execute(ctx, req.Id, req.Cash)

		if err != nil {
			switch {
			case errors.Is(err, account.ErrInvalidAccountId),
				errors.Is(err, account.ErrInvalidAccountCash):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
