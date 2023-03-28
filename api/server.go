package api

import (
	"github.com/gin-gonic/gin"
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

	// Add the routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router

	return server
}
