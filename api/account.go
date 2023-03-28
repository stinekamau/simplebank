package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/stinekamau/simplebank/db/sqlc"
)

type AccountRequestParams struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func (s *Server) createAccount(c *gin.Context) {
	var req AccountRequestParams

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error creating accounts, %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}
	// Insert a new account to the database
	account, err := s.store.CreateAccount(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)

}
