package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/treatment"
	"medical-record-api/pkg/middleware"
)

func RegisterTreatmentRoutes(rg *gin.RouterGroup, handler *treatment.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.AddTreatment)
	rg.GET("", handler.GetTreatments)
	rg.DELETE("/:id", handler.DeleteTreatment)
	rg.PUT("/:id", handler.UpdateTreatment)
}
