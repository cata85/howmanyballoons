package api

import (
	"net/http"

	"github.com/cata85/balloons/db"
	"github.com/gin-gonic/gin"
)

/**
 * Method:   GET
 * Endpoint: /all
 * Display's all the saved balloon objects.
 */
func HandlerGetAll(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	if session.Values["is_admin"] == true { // The compiler didn't like session.Values["is_admin"]
		var balloonObjects, _ = db.GetAllBalloonObjects()

		c.HTML(http.StatusOK, "all.html", gin.H{
			"balloonObjects": balloonObjects,
		})
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

/**
 * METHOD: GET
 * Endpoint: /all/delete/:id
 * This soft deletes a balloon object and returns the user back to the /all page.
 * Notes: I felt like using a GET request. That is all. Was doing some hacky html/js workarounds
 *        to maintain the User's page view on the /all endpoint without them having to see any url
 *        changes (like http://xxxxxxxx/all/delete?id=12) in their browser as well as ensuring page refresh
 *        worked smoothly with the updates.
 */
func HandlerDeleteAllSingle(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	if session.Values["is_admin"] == true { // The compiler didn't like session.Values["is_admin"]
		id := c.Param("id")
		db.DeleteSingle(id)
		var balloonObjects, _ = db.GetAllBalloonObjects()

		c.HTML(http.StatusOK, "all.html", gin.H{
			"balloonObjects": balloonObjects,
		})
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}
