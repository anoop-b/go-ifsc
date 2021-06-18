package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ifsc/helpers"
	"go-ifsc/helpers/token"
	"go-ifsc/models"
	"net/http"
)

func GetAuth(c *gin.Context) {
	var auth models.User
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusForbidden, "unauthorised request")
		return
	}

	ck := helpers.NewCacheServer()

	cacheUser, uExists := ck.GetCache("username")
	cachePassword, pExists := ck.GetCache("password")

	if !uExists || !pExists {
		c.String(http.StatusUnauthorized, "unauthorised, please register")
		c.Abort()
		return
	}

	if cacheUser.(string) != auth.Username && cachePassword.(string) != auth.Password {
		c.String(http.StatusUnauthorized, "unauthorised, please register")
		c.Abort()
		return
	}

	jwt := token.NewJWTService()
	jwtToken := jwt.GenerateToken(auth.Username, auth.Password)
	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})

}
