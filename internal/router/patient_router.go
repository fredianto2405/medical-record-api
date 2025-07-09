package router

import (
	"github.com/gin-gonic/gin"
	patient "medical-record-api/internal/patient/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterPatientRoutes(rg *gin.RouterGroup, handler *patient.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetPatients)
	rg.POST("", handler.AddPatient)
	rg.PUT("/:id", handler.UpdatePatient)
	rg.DELETE("/:id", handler.DeletePatient)
}
