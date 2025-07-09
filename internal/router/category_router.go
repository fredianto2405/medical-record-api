package router

import (
	"github.com/gin-gonic/gin"
	medicine "medical-record-api/internal/medicine/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, handler *medicine.CategoryHandler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("/categories", handler.GetCategories)
	rg.POST("/categories", handler.AddCategory)
	rg.PUT("/categories/:id", handler.UpdateCategory)
	rg.DELETE("/categories/:id", handler.DeleteCategory)
}
