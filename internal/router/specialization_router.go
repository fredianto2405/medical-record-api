package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/specialization"
	"medical-record-api/pkg/middleware"
)

func RegisterSpecializationRoutes(rg *gin.RouterGroup, handler *specialization.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetSpecializations)
	rg.POST("", handler.AddSpecialization)
	rg.PUT("/:id", handler.UpdateSpecialization)
	rg.DELETE("/:id", handler.DeleteSpecialization)
}
