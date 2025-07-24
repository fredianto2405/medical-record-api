package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/auth"
	"medical-record-api/pkg/middleware"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	rg.POST("/login", handler.Login)
	rg.POST("/change-password", middleware.JWTAuthMiddleware(), handler.ChangePassword)
}
