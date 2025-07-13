package router

import (
	"github.com/gin-gonic/gin"
	payment "medical-record-api/internal/payment/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterInsuranceRoutes(rg *gin.RouterGroup, handler *payment.InsuranceHandler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("", handler.GetInsurances)
	rg.POST("", handler.AddInsurance)
	rg.PUT("/:id", handler.UpdateInsurance)
	rg.DELETE("/:id", handler.DeleteInsurance)
}
