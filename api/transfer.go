package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/stinekamau/simplebank/db/sqlc"
)

type TransferRequestParams struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (s *Server) createTransfer(c *gin.Context) {
	var req TransferRequestParams

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error creating transfer, %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !s.validateAccount(c, req.FromAccountID, req.Currency) {
		return
	}

	if !s.validateAccount(c, req.ToAccountID, req.Currency) {
		return
	}

	args := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	account, err := s.store.TransferTx(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)

}

func (s *Server) validateAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := s.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Account does not exist"})
			return false
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating account"})
		return false
	}
	if strings.ToUpper(account.Currency) != currency {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Different currencies"})
		return false

	}
	return true

}
