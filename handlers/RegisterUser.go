package handlers

import (
	"github.com/gin-gonic/gin"
	"go-ifsc/helpers"
	"go-ifsc/models"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusUnauthorized, "validation failed, minimum 5 characters for username and 12 characters for password")
		return
	}
	cache := helpers.NewCacheServer()
	cache.SetCache("username", user.Username)
	cache.SetCache("password", user.Password)
	c.String(200, "successfully registered")
}
