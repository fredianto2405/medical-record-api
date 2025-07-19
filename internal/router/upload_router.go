package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/upload"
	"medical-record-api/pkg/middleware"
)

func RegisterUploadRoutes(rg *gin.RouterGroup, handler *upload.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.Upload)
}
