package accounts

import (
	"errors"
	"net/http"

	mooc "github.com/Metalisaac321/stock-market-simulator/internal"
	"github.com/Metalisaac321/stock-market-simulator/internal/creating"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Id   string `json:"id"`
	Cash uint   `json:"cash"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(accountService creating.AccountService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := accountService.CreateAccount(ctx, req.Id, req.Cash)

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidAccountId),
				errors.Is(err, mooc.ErrInvalidAccountCash):
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
