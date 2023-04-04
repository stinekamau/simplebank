package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/stinekamau/simplebank/db/sqlc"
)

type Server struct {
	// need the store  to access the db
	store *db.Store
	// router to map requests to handler functions
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	// Create a new store
	server := &Server{store: store}

	// add the engine
	router := gin.Default()

	// Check for the validator engine
	engine := binding.Validator.Engine().(validator.Validate)

	// Register the validation function
	engine.RegisterValidation("currency", validCurrency)

	// Add the routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/delete", server.deleteAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router

	return server
}
