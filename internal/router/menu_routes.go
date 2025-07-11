package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/menu"
	"medical-record-api/pkg/middleware"
)

func RegisterMenuRoutes(rg *gin.RouterGroup, handler *menu.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetByRoleID)
}
