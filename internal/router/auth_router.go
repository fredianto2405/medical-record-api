package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/auth"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	rg.POST("/login", handler.Login)
}
