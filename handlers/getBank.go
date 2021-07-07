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

var banksJson map[string]models.Bank
var bankNamesJson map[string]string

//GetBank is the default handler func for fetching IFSC
func GetBank(directory *embed.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		ifsc := c.MustGet("sanitisedIFSC").(string)
		bankCode := ifsc[0:4]

		bankNamesBlob, _ := directory.ReadFile("Data/" + "banknames.json")
		err := json.Unmarshal(bankNamesBlob, &bankNamesJson)
		if err != nil {
			log.Fatalln("Failed to unmarshal", err)
		}

		_, exists := bankNamesJson[bankCode]
		if !exists {
			c.AbortWithStatusJSON(http.StatusNotFound, "Not Found")
			return
		}

		jsonObject, err := directory.ReadFile("Data/" + bankCode + ".json")
		if os.IsNotExist(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, "Not Found")
			return
		}

		err = json.Unmarshal(jsonObject, &banksJson)
		if err != nil {
			log.Fatalln("Failed to unmarshal", err)
		}

		if ifsc != "" {
			payload, exists := banksJson[ifsc]
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
