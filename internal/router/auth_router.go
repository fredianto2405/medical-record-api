package router

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/auth"
	"medical-record-api/pkg/middleware"
	"time"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	loginLimiter := middleware.RateLimiterMiddleware(5, time.Minute)
	rg.POST("/login", loginLimiter, handler.Login)

	rg.POST("/change-password", middleware.JWTAuthMiddleware(), handler.ChangePassword)

	forgotLimiter := middleware.RateLimiterMiddleware(3, 10*time.Minute)
	rg.POST("/forgot-password", forgotLimiter, handler.ForgotPassword)

	resetLimiter := middleware.RateLimiterMiddleware(3, time.Minute)
	rg.POST("/reset-password/:token", resetLimiter, handler.ResetPassword)
}
