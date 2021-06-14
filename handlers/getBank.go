package handlers

import (
	"encoding/json"
	"go-ifsc/helpers"
	"go-ifsc/models"

	"github.com/gin-gonic/gin"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//GetBank is the default handler func for fetching IFSC
func GetBank(c *gin.Context) {
	var banks map[string]models.Bank
	ifsc := c.MustGet("sanitisedIFSC").(string)
	bankCode := ifsc[0:4]

	jsonFilePath := filepath.Join("Data", bankCode+".json")

	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusNotFound, "Not Found")
		return
	}
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatalln("error opening json file", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &banks)
	if err != nil {
		log.Fatalln("Failed to unmarshal", err)
	}

	if ifsc != "" {
		payload, exists := banks[ifsc]
		if exists {
			c.SecureJSON(http.StatusOK, payload)
			helpers.SetCache(ifsc, payload)
			return
		}
		c.AbortWithStatusJSON(http.StatusNotFound, "Not Found")
	}
}
