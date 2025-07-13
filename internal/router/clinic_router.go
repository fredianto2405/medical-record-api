package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/clinic"
	"medical-record-api/pkg/middleware"
)

func RegisterClinicRoutes(rg *gin.RouterGroup, handler *clinic.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetClinic)
	rg.PUT("/:id", handler.UpdateClinic)
}
