package api

import (
	"net/http"

	"github.com/cata85/balloons/db"
	"github.com/cata85/balloons/types"
	"github.com/gin-gonic/gin"
)

var balloonObject *(types.BalloonObject)

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

	c.HTML(http.StatusOK, "index.html", gin.H{
		"itemName":       balloonObject.Name,
		"itemWeight":     balloonObject.Weight,
		"itemBalloons":   balloonObject.Balloons,
		"itemWeightType": balloonObject.WeightType,
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

	balloonObject.Name = c.PostForm("itemName")
	balloonObject.Weight = c.PostForm("itemWeight")
	balloonObject.WeightType = c.PostForm("itemWeightType")
	if balloonObject.Balloons != "" {
		balloonObject.Balloons = Calculate(balloonObject.Weight, balloonObject.WeightType)
	}
	db.SaveBalloonObject(*balloonObject)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"itemName":       balloonObject.Name,
		"itemWeight":     balloonObject.Weight,
		"itemBalloons":   balloonObject.Balloons,
		"itemWeightType": balloonObject.WeightType,
	})
}
