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

//Bank struct denotes bank response struct
type Bank struct {
	BANK     string `json:"BANK"`
	IFSC     string `json:"IFSC"`
	BRANCH   string `json:"BRANCH"`
	CENTRE   string `json:"CENTRE"`
	DISTRICT string `json:"DISTRICT"`
	STATE    string `json:"STATE"`
	ADDRESS  string `json:"ADDRESS"`
	CONTACT  string `json:"CONTACT"`
	IMPS     bool   `json:"IMPS"`
	CITY     string `json:"CITY"`
	UPI      bool   `json:"UPI"`
	MICR     string `json:"MICR"`
	NEFT     bool   `json:"NEFT"`
	RTGS     bool   `json:"RTGS"`
}

//GetBank is the default handler func for fetching IFSC
func GetBank(c *gin.Context) {
	var banks map[string]Bank
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
