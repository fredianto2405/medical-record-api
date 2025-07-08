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
		if authHeader == "" {
			response.Respond(c, http.StatusUnauthorized, false, constant.MsgAuthHeaderMissing, nil, nil)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Respond(c, http.StatusUnauthorized, false, constant.MsgInvalidAuthHeader, nil, nil)
			c.Abort()
			return
		}

		claims, err := jwt.ValidateJWT(parts[1])
		if err != nil {
			response.Respond(c, http.StatusUnauthorized, false, constant.MsgInvalidToken, nil, nil)
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
