package handlers

import (
	"encoding/json"
	"go-ifsc/helpers"

	"github.com/gin-gonic/gin"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GetBank(c *gin.Context) {
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
	var banks map[string]interface{}
	err = json.Unmarshal(byteValue, &banks)
	if err != nil {
		log.Fatalln("Failed to unmarshal", err)
	}

	if ifsc != "" {
		if banks[ifsc] != nil {
			c.SecureJSON(http.StatusOK, banks[ifsc])
			helpers.SetCache(ifsc, banks[ifsc])
			return
		}
		c.AbortWithStatusJSON(http.StatusNotFound, "Not Found")
	}
}
