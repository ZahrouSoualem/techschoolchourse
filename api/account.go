package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "tutorial.sqlc.dev/app/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required" `
	Currency string `json:"currency" binding:"required" `
}

func (server *Server) createAccount(ctx *gin.Context) {

	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResqponse(err))
		return
	}

	args := db.CreateAuthorParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAuthor(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResqponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type getAccountParams struct {
	ID int64 `uri:"id" binding:"required,min=1" `
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResqponse(err))
		return
	}

	account, err := server.store.GetAuthor(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResqponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResqponse(err))
		return
	}

	//account := &db.Account{}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResqponse(err))
		return
	}
	fmt.Println((req.PageID - 1) * req.PageSize)

	args := db.ListAuthorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAuthors(ctx, args)
	fmt.Println("Zahrou")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResqponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
