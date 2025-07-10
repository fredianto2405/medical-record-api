package router

import (
	"github.com/gin-gonic/gin"
	patient "medical-record-api/internal/patient/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterInsurancePatientRoutes(rg *gin.RouterGroup, handler *patient.InsurancePatientHandler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("/insurance", handler.AddInsurancePatient)
	rg.DELETE("/insurance/:patientID", handler.DeleteInsurancePatient)
}
