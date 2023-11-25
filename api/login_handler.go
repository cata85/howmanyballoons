package api

import (
	"log"
	"net/http"

	"github.com/cata85/balloons/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var user *(types.User)

/**
 * Method:   POST
 * Endpoint: /
 * When the user submits the balloon object creation form.
 * Sends the balloon object to be calculated and upserts into postgres
 */
func HandlerPostLogin(c *gin.Context) {
	if user == nil {
		user = new(types.User)
	}

	user.Name = c.PostForm("username")
	hash, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), 0)
	user.Password = string(hash)
	log.Printf("Password: %v\n", user.Password)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"itemName":       balloonObject.Name,
		"itemWeight":     balloonObject.Weight,
		"itemBalloons":   balloonObject.Balloons,
		"itemWeightType": balloonObject.WeightType,
	})
}
