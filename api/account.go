package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "tutorial.sqlc.dev/app/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required" `
	Currency string `json:"currency" binding:"required, oneof:USD EUR" `
}

func (server *Server) createAccount(ctx *gin.Context) {

	var req createAccountRequest
	fmt.Println("Hello World")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResqponse(err))
		return
	}
	fmt.Println("Hello World 2")
	fmt.Println(req.Owner)
	args := db.CreateAuthorParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	fmt.Println("Hello World 3")
	fmt.Println(req.Owner)
	account, err := server.store.CreateAuthor(ctx, args)

	fmt.Println("Hello World 2")
	fmt.Println(account.Currency)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResqponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
