package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginLimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"
	"time"
)

func RateLimiterMiddleware(limit int, per time.Duration) gin.HandlerFunc {
	rate := limiter.Rate{
		Period: per,
		Limit:  int64(limit),
	}

	store := memoryStore.NewStore()
	instance := limiter.New(store, rate)

	return ginLimiter.NewMiddleware(instance)
}
