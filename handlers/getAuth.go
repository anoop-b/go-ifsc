package handlers

import (
	"github.com/gin-gonic/gin"
	"go-ifsc/helpers"
	"go-ifsc/helpers/token"
	"go-ifsc/models"
	"log"
	"net/http"
)

func GetAuth(c *gin.Context) {
	var auth models.User
	err := c.ShouldBindJSON(&auth)
	if err != nil {
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

	paseto, err := token.NewPasetoMaker()
	if err != nil {
		log.Panic(err.Error())
	}
	pasetoToken := paseto.GenerateToken(auth.Username, auth.Password)
	c.SetCookie(
		"token",
		pasetoToken,
		int(token.MaxAge.Seconds()),
		"/",
		"localhost",
		false,
		true,
	)
	c.String(200, "Login Successful")
}
