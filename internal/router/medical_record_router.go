package router

import (
	"github.com/gin-gonic/gin"
	handler "medical-record-api/internal/medical_record/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterMedicalRecordRoutes(rg *gin.RouterGroup, handler *handler.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("/status", handler.GetStatuses)
	rg.POST("", handler.AddMedicalRecord)
	rg.DELETE("/:id", handler.DeleteMedicalRecord)
	rg.DELETE("/:id/nurse/:nurseId", handler.DeleteNurseAssignment)
	rg.DELETE("/:id/treatment/:treatmentId", handler.DeleteTreatmentDetail)
	rg.DELETE("/:id/medicine/:medicineId", handler.DeleteRecipe)
}
