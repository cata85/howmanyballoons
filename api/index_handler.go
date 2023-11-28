package api

import (
	"net/http"

	"github.com/cata85/balloons/db"
	"github.com/cata85/balloons/types"
	"github.com/gin-gonic/gin"
)

/**
 * Method:   GET
 * Endpoint: /
 * Default homepage for the User.
 */
func HandlerGetIndex(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	if session.Values["name"] == nil {
		session.Values["name"] = ""
		session.Values["logged_in"] = false
		session.Values["is_admin"] = false
		session.Values["item_name"] = ""
		session.Values["item_weight"] = ""
		session.Values["item_balloons"] = ""
		session.Values["weight_type"] = "Pound"
		_ = session.Save(c.Request, c.Writer)
	}
	savedBalloonObjects, _ := db.GetAllActiveBalloonObjects()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"itemName":            session.Values["item_name"],
		"itemWeight":          session.Values["item_weight"],
		"itemBalloons":        session.Values["item_balloons"],
		"itemWeightType":      session.Values["weight_type"],
		"savedBalloonObjects": savedBalloonObjects,
		"name":                session.Values["name"],
		"is_admin":            session.Values["is_admin"],
		"logged_in":           session.Values["logged_in"],
	})
}

/**
 * Method:   POST
 * Endpoint: /
 * When the user submits the balloon object creation form.
 * Sends the balloon object to be calculated and upserts into postgres
 */
func HandlerPostIndex(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	session.Values["item_name"] = c.PostForm("itemName")
	session.Values["item_weight"] = c.PostForm("itemWeight")
	session.Values["weight_type"] = c.PostForm("itemWeightType")
	balloons := Calculate(c.PostForm("itemWeight"), c.PostForm("itemWeightType"))
	session.Values["item_balloons"] = balloons
	_ = session.Save(c.Request, c.Writer)
	balloonObject := types.BalloonObject{
		Name:       c.PostForm("itemName"),
		Weight:     c.PostForm("itemWeight"),
		Balloons:   balloons,
		WeightType: c.PostForm("itemWeightType"),
	}
	db.SaveBalloonObject(balloonObject)
	savedBalloonObjects, _ := db.GetAllActiveBalloonObjects()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"itemName":            session.Values["item_name"],
		"itemWeight":          session.Values["item_weight"],
		"itemBalloons":        session.Values["item_balloons"],
		"itemWeightType":      session.Values["weight_type"],
		"savedBalloonObjects": savedBalloonObjects,
		"name":                session.Values["name"],
		"is_admin":            session.Values["is_admin"],
		"logged_in":           session.Values["logged_in"],
	})
}
