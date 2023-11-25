package api

import (
	"net/http"
	"os"

	"github.com/cata85/balloons/db"
	helper "github.com/cata85/balloons/helpers"
	"github.com/cata85/balloons/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var balloonObject *types.BalloonObject
var store *sessions.CookieStore

/**
 * Method:   GET
 * Endpoint: /
 * Default homepage for the User.
 */
func HandlerGetIndex(c *gin.Context) {
	if balloonObject == nil {
		balloonObject = new(types.BalloonObject)
		balloonObject.WeightType = "Pound"
	}
	if store == nil {
		config := helper.Config()
		gorillaConfig := config["gorilla"]
		store = sessions.NewCookieStore([]byte(os.Getenv(helper.String(helper.Get(gorillaConfig, "key")))))
	}

	session, _ := store.Get(c.Request, "session")
	if session.Values["name"] == nil {
		session.Values["name"] = ""
		session.Values["logged_in"] = false
		session.Values["is_admin"] = false
		_ = session.Save(c.Request, c.Writer)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"itemName":       balloonObject.Name,
		"itemWeight":     balloonObject.Weight,
		"itemBalloons":   balloonObject.Balloons,
		"itemWeightType": balloonObject.WeightType,
		"name":           session.Values["name"],
		"is_admin":       session.Values["is_admin"],
		"logged_in":      session.Values["logged_in"],
	})
}

/**
 * Method:   POST
 * Endpoint: /
 * When the user submits the balloon object creation form.
 * Sends the balloon object to be calculated and upserts into postgres
 */
func HandlerPostIndex(c *gin.Context) {
	if balloonObject == nil {
		balloonObject = new(types.BalloonObject)
		balloonObject.WeightType = "Pound"
	}
	session, _ := store.Get(c.Request, "session")

	balloonObject.Name = c.PostForm("itemName")
	balloonObject.Weight = c.PostForm("itemWeight")
	balloonObject.WeightType = c.PostForm("itemWeightType")
	balloonObject.Balloons = Calculate(balloonObject.Weight, balloonObject.WeightType)
	db.SaveBalloonObject(*balloonObject)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"itemName":       balloonObject.Name,
		"itemWeight":     balloonObject.Weight,
		"itemBalloons":   balloonObject.Balloons,
		"itemWeightType": balloonObject.WeightType,
		"name":           session.Values["name"],
		"is_admin":       session.Values["is_admin"],
		"logged_in":      session.Values["logged_in"],
	})
}
