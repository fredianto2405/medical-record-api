package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/doctor"
	"medical-record-api/pkg/middleware"
)

func RegisterDoctorRoutes(rg *gin.RouterGroup, handler *doctor.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetDoctors)
	rg.POST("", handler.AddDoctor)
	rg.PUT("/:id", handler.UpdateDoctor)
	rg.DELETE("/:id", handler.DeleteDoctor)
}
