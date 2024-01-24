package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/di"
)

func HealthProblemRoutes(conn *sql.DB, r *gin.RouterGroup) {

	healthProblemHandler := di.ConfigHealthProblemDI(conn)

	healthProblemGroup := r.Group("/health-problems")

	healthProblemGroup.POST("/:clientId", healthProblemHandler.CreateHealthProblem)
	healthProblemGroup.PATCH("/:id", healthProblemHandler.UpdateHealthProblem)

}
