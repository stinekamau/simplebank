package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/stinekamau/simplebank/db/sqlc"
	"github.com/stinekamau/simplebank/utils"
)

type CreateUserRequestParams struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string       `json:"username" binding:"required,alphanum"`
	Password          string       `json:"password" binding:"required,min=6"`
	FullName          string       `json:"full_name" binding:"required"`
	Email             string       `json:"email" binding:"required,email"`
	CreatedAt         time.Time    `json:"created_at"`
	PasswordChangedAt sql.NullTime `json:"password_changed_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
	}

}

// func (s *Server) Start(address string) error {
// 	return s.router.Run(address)
// }

func (s *Server) createUser(c *gin.Context) {
	var req CreateUserRequestParams

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error creating accounts, %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		FullName:       req.FullName,
	}
	// Insert a new account to the database
	account, err := s.store.CreateUser(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user_response"`
}

func (s *Server) loginUser(c *gin.Context) {
	var req loginUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Retrieve the user request from the database
	user, err := s.store.GetUser(c, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

	}

	// user found
	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
	}

	token, err := s.tokenMaker.CreateToken(user.Username, s.config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	resp := loginUserResponse{
		AccessToken: token,
		User:        newUserResponse(user),
	}

	c.JSON(http.StatusOK, resp)

}
