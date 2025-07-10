package router

import (
	"github.com/gin-gonic/gin"
	patient "medical-record-api/internal/patient/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterEmergencyContactRoutes(rg *gin.RouterGroup, handler *patient.EmergencyContactHandler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("/emergency-contact", handler.AddEmergencyContact)
	rg.DELETE("/emergency-contact/:patientID", handler.DeleteEmergencyContact)
}
