package middlewares

import (

	"github.com/gin-gonic/gin"
	"go-ifsc/helpers"

	"log"
	"net/http"
)

func CacheCheck(c *gin.Context) {
	ifsc, err := helpers.ValidIfsc(c.Param("ifsc"))
	c.Set("sanitisedIFSC", ifsc)
	if !err {
		c.AbortWithStatusJSON(http.StatusNotFound, "Not a Valid IFSC")
	}
	response, exists := helpers.GetCache(ifsc)
	if exists {
		log.Println("cache hit")
		c.SecureJSON(http.StatusOK, response)
		c.Abort()
	} else {
		c.Next()
	}
}
