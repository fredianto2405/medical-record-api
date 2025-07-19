package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/user"
	"medical-record-api/pkg/middleware"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *user.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.AddUser)
	rg.DELETE("/:id", handler.DeleteUser)
}
