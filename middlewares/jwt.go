package middlewares

import (
	"go-ifsc/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer "

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := authHeader[len(BEARER_SCHEMA):]
			jwt := helpers.NewJWTService()
			token, err := jwt.ValidateToken(tokenString)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
				return
			}
			if token.Valid {
				c.Next()
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
		}
		c.Abort()
	}
}
