package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/nurse"
	"medical-record-api/pkg/middleware"
)

func RegisterNurseRoutes(rg *gin.RouterGroup, handler *nurse.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetNurses)
	rg.POST("", handler.AddNurse)
	rg.PUT("/:id", handler.UpdateNurse)
	rg.DELETE("/:id", handler.DeleteNurse)
}
