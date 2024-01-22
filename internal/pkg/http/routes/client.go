package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/di"
)

func ClientRoutes(conn *sql.DB, r *gin.RouterGroup) {

	clientHandler := di.ConfigClientDI(conn)

	clientGroup := r.Group("/clients")

	clientGroup.POST("", clientHandler.CreateClient)
	clientGroup.GET("", clientHandler.FindAllClients)
	clientGroup.GET("/:id", clientHandler.FindClientById)

}
