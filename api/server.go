package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sajitron/travel-agency/db/sqlc"
	"github.com/sajitron/travel-agency/util"
)

type Server struct {
	config util.Config
	router *gin.Engine
	store  db.Store
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/cheers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "say cheese")
	})

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
