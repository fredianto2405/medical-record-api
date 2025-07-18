package router

import (
	"github.com/gin-gonic/gin"
	handler "medical-record-api/internal/medical_record/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterMedicalRecordRoutes(rg *gin.RouterGroup, handler *handler.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("/status", handler.GetStatuses)

	rg.GET("", handler.GetMedicalRecords)
	rg.POST("", handler.AddMedicalRecord)
	rg.PUT("/:id", handler.UpdateMedicalRecord)
	rg.PUT("/:id/status", handler.UpdateStatusMedicalRecord)
	rg.DELETE("/:id", handler.DeleteMedicalRecord)

	rg.POST("/:id/nurse", handler.AddNurseAssignment)
	rg.DELETE("/:id/nurse/:nurseId", handler.DeleteNurseAssignment)

	rg.POST("/:id/treatment", handler.AddTreatmentDetail)
	rg.DELETE("/:id/treatment/:treatmentId", handler.DeleteTreatmentDetail)

	rg.POST("/:id/medicine", handler.AddRecipe)
	rg.DELETE("/:id/medicine/:medicineId", handler.DeleteRecipe)
}
