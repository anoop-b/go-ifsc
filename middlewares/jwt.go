package middlewares

import (
	"go-ifsc/helpers/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer "

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := authHeader[len(BEARER_SCHEMA):]
			jwt := token.NewJWTService()
			jwtToken, err := jwt.ValidateToken(tokenString)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
				return
			}
			if jwtToken.Valid {
				c.Next()
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
		}
		c.Abort()
	}
}
