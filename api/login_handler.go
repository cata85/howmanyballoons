package api

import (
	"net/http"

	"github.com/cata85/balloons/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/**
 * Method:   POST
 * Endpoint: /login
 * When the user submits username and password to login.
 */
func HandlerPostLogin(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	user.Name = c.PostForm("username")
	user.Password = c.PostForm("password")
	dbUser := db.GetOneUser(user.Name)
	if dbUser != nil {
		err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
		if err == nil {
			if user.Name == "admin" {
				session.Values["is_admin"] = true
			} else {
				session.Values["is_admin"] = false
			}
			session.Values["name"] = user.Name
			session.Values["logged_in"] = true

			_ = session.Save(c.Request, c.Writer)
		}
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
 * Endpoint: /signup
 * When the user submits username and password to signup.
 */
func HandlerPostSignup(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	user.Name = c.PostForm("username")
	hash, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	user.Password = string(hash)
	db.SaveUser(*user)
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
 * Method:   GET
 * Endpoint: /logout
 * Logs the user out.
 */
func HandlerGetLogout(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	session.Values["name"] = ""
	session.Values["logged_in"] = false
	session.Values["is_admin"] = false
	_ = session.Save(c.Request, c.Writer)

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
