package api

import (
	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server to serve HTTP requests for bank services
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer instance to setup/create a new HTTP/API server routing for services on the server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		val.RegisterValidation("currency", validCurrency)
	}


	// Add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	router.POST("/transfers", server.createTransfer)






	server.router = router
	return server
}

// handling http server 
// Start runs the HTTP server on a specific address...to listen to API requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}	

// implementing errorResponse from account.go
func errorReponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}