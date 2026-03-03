package router

import (
	"github.com/gin-gonic/gin"
	"github.com/theozdev/signa_q/internal/healthcheck"
)

func InitRouters(r *gin.RouterGroup) {
	healthGroup := r.Group("/healthcheck")
	healthcheck.InitRoutes(healthGroup)
}
