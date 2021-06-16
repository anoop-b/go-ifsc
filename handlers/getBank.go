package handlers

import (
	"embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-ifsc/helpers"
	"go-ifsc/models"

	"log"
	"net/http"
	"os"
)

//GetBank is the default handler func for fetching IFSC
func GetBank(directory *embed.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		var banks map[string]models.Bank
		ifsc := c.MustGet("sanitisedIFSC").(string)
		bankCode := ifsc[0:4]
		jsonObject, err := directory.ReadFile("Data/" + bankCode + ".json")
		if os.IsNotExist(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, "Not Found")
			return
		}
		err = json.Unmarshal(jsonObject, &banks)
		if err != nil {
			log.Fatalln("Failed to unmarshal", err)
		}

		if ifsc != "" {
			payload, exists := banks[ifsc]
			if exists {
				c.SecureJSON(http.StatusOK, payload)
				ck := helpers.NewCacheServer()
				ck.SetCache(ifsc, payload)
				return
			}
			c.AbortWithStatusJSON(http.StatusNotFound, "Not Found")
		}
	}
}
