package router

import (
	"health-service/handler"

	"github.com/gin-gonic/gin"
)

func AttachRouter(router *gin.Engine) {
	router.GET("/api/health", handler.HealthHandler)
}
