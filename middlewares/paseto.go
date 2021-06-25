package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-ifsc/helpers/token"
	"log"
	"net/http"
)

func AuthorizePaseto() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader, err := c.Request.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
			return
		}
		tokenString := authHeader.Value
		paseto, err := token.NewPasetoMaker()
		if err != nil {
			log.Panic(err)
		}
		_, err = paseto.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
			return
		}
	}
}
