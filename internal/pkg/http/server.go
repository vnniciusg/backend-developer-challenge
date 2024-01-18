package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/routes"
)

type Server struct {
	HttpPort string
	Router   *gin.Engine
	DB       *sql.DB
}

func NewServer(httpPort string, r *gin.Engine, conn *sql.DB) *Server {
	server := &Server{
		HttpPort: httpPort,
		Router:   r,
		DB:       conn,
	}

	server.InitRoutes()

	return server

}

func (s *Server) InitRoutes() {
	v1 := s.Router.Group("/api/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		restSuccess := responseshttp.NewOk("pong", nil)
		ctx.JSON(restSuccess.Code, restSuccess)
		return
	})

	routes.ClientRoutes(s.DB, v1)
}

func (s *Server) Run() error {
	return s.Router.Run(s.HttpPort)
}
