package api

import (
	db "github.com/Kawaeugtkp/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	// interfaceは*をつけなくていいらしい（？）
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// POSTの定義を見に行くと"POST(relativePath string, handlers ...HandlerFunc)"となっていて、
	// HandlerFuncはtype HandlerFunc func(*Context)だから、createAccount(ctx *gin.Context)のような定義をしたものが引数として必要
	router.POST("/accounts", server.createAccount) 
	router.GET("/accounts/:id", server.getAccount) 
	router.GET("/accounts", server.listAccount) 

	router.POST("/transfers", server.createTransfer) 

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}