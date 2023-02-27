package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "tutorial.sqlc.dev/app/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account" binding:"required,min=1" `
	ToAccountID   int64  `json:"to_account" binding:"required,min=1"`
	Amount        int64  `json:"from_account" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {

	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResqponse(err))
		return
	}
	_, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}
	_, valid = server.validAccount(ctx, req.FromAccountID, req.Currency)

	if !valid {
		return
	}

	args := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResqponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAuthor(ctx, accountID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResqponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResqponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResqponse(err))
		return account, false
	}

	return account, true
}

func errorResponse(err error) {
	panic("unimplemented")
}
