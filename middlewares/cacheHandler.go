package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-ifsc/helpers"

	"log"
	"net/http"
)

func CacheCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		ifsc, err := helpers.ValidIfsc(c.Param("ifsc"))
		if !err {
			c.AbortWithStatusJSON(http.StatusNotFound, "Not a Valid IFSC")
			return
		}
		c.Set("sanitisedIFSC", ifsc)
		response, exists := helpers.NewCacheServer().GetCache(ifsc)
		if exists {
			log.Println("cache hit")
			c.SecureJSON(http.StatusOK, response)
			c.Abort()
			return
		}
	}
}
