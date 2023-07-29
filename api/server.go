package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/db/sqlc"
)

// Server serves HTTP requests for our simple bank
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates http routes
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// register a validator that can be used in Gin struct tag bindings
	// meaning that validCurrency will validate a field value when there's a binding:"currency" in its struct tag
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
