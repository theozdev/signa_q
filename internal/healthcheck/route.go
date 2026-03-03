package healthcheck

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	handle := NewHandler()
	r.GET("", handle.Ping)
}
