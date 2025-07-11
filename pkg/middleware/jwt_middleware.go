package middleware

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	"medical-record-api/pkg/jwt"
	"medical-record-api/pkg/response"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Respond(c, http.StatusUnauthorized, false, constant.MsgAuthHeaderMissing, nil, nil)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateJWT(tokenStr)
		if err != nil {
			response.Respond(c, http.StatusUnauthorized, false, err.Error(), nil, nil)
			c.Abort()
			return
		}

		// Inject claims into context
		c.Set("user", claims)
		c.Next()
	}
}
