package router

import (
	"github.com/gin-gonic/gin"
	medicine "medical-record-api/internal/medicine/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterUnitRoutes(rg *gin.RouterGroup, handler *medicine.UnitHandler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("/units", handler.GetUnits)
	rg.POST("/units", handler.AddUnit)
	rg.PUT("/units/:id", handler.UpdateUnit)
	rg.DELETE("/units/:id", handler.DeleteUnit)
}
