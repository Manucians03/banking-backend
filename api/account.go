package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/Manucians03/banking-backend/db/sqlc"
	"github.com/Manucians03/banking-backend/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required, currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, account)
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	authoPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authoPayload.Username {
		err := errors.New("account does not belong to the authenticated user")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
