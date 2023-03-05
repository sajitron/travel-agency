package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sajitron/travel-agency/db/sqlc"
	"github.com/sajitron/travel-agency/token"
	"github.com/sajitron/travel-agency/util"
)

type Server struct {
	config     util.Config
	router     *gin.Engine
	tokenMaker token.Maker
	store      db.Store
}

// NewServer creates a new server and sets up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("unable to initialise token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
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
