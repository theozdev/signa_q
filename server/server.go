package server

import (
	"github.com/gin-gonic/gin"
	"github.com/theozdev/signa_q/internal/router"
)

func Init() {
	port := "3000"

	r := gin.Default()
	api := r.Group("/api")

	router.InitRouters(api)

	if port == "" {
		port = "3000"
	}
	r.SetTrustedProxies(nil)

	r.Run(":" + port)
}
