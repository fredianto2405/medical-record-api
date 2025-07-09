package router

import (
	"github.com/gin-gonic/gin"
	medicine "medical-record-api/internal/medicine/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterMedicineRoutes(rg *gin.RouterGroup, handler *medicine.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetMedicines)
	rg.POST("", handler.AddMedicine)
	rg.PUT("/:id", handler.UpdateMedicine)
	rg.DELETE("/:id", handler.DeleteMedicine)
}
