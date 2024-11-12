package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/vivek-344/banking-system/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `binding:"required" json:"owner"`
	Currency string `binding:"required,oneof=USD EUR" json:"currency" `
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

type getAccountRequest struct {
	ID int64 `binding:"required,min=1" uri:"id"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	PageID   int32 `binding:"required,min=1" form:"page_id"`
	PageSize int32 `binding:"required,min=5,max=10" form:"page_size"`
}

type listAccountsResponse struct {
	Accounts []db.Account `json:"accounts"`
	Count    int32        `json:"count"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listAccountsResponse{
		Accounts: accounts,
		Count:    int32(len(accounts)),
	})
}

type updateAccountURI struct {
	ID int64 `binding:"required,min=1" uri:"id"`
}

type updateAccountParams struct {
	Balance *int64 `binding:"min=0" json:"balance"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var reqURI updateAccountURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqBody updateAccountParams
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if reqBody.Balance == nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("field validation for 'balance' failed: 'balance' is required")))
		return
	}
	if *reqBody.Balance < 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("field validation for 'balance' failed: minimum value is 0")))
		return
	}

	account, err := server.store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:      reqURI.ID,
		Balance: *reqBody.Balance,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type deleteAccountRequest struct {
	ID int64 `binding:"required,min=1" uri:"id"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteAccount(ctx, req.ID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "account deleted successfully"})
}
