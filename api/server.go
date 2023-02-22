package api

import (
	"github.com/gin-gonic/gin"
	db "tutorial.sqlc.dev/app/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	// add route to this router

	router.POST("/accounts", server.createAccount)
	server.router = router

	return server

}

// run HTTP request on a specific address
func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errorResqponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}