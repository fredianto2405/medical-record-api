package router

import (
	"github.com/gin-gonic/gin"
	payment "medical-record-api/internal/payment/handler"
	"medical-record-api/pkg/middleware"
)

func RegisterPaymentRoutes(rg *gin.RouterGroup, handler *payment.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.GET("/methods", handler.GetMethods)
	rg.POST("/methods", handler.AddMethod)
	rg.PUT("/methods/:id", handler.UpdateMethod)
	rg.DELETE("/methods/:id", handler.DeleteMethod)
	rg.GET("/statuses", handler.GetStatuses)
}
