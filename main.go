package main

import (
	"log"

	"github.com/cata85/balloons/api"
	"github.com/cata85/balloons/db"
	helper "github.com/cata85/balloons/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	config := helper.Config()
	serverConfig := config["server"]

	db.InitializeTables()
	api.Initialize()
	router := gin.Default()
	router.LoadHTMLGlob("html/templates/*/*.html")
	router.Static("/html/static", "./html/static")

	router.GET("/", api.HandlerGetIndex)
	router.POST("/", api.HandlerPostIndex)
	router.GET("/all", api.HandlerGetAll)
	router.GET("/all/delete/:id", api.HandlerDeleteAllSingle)
	router.POST("/login", api.HandlerPostLogin)
	router.POST("/signup", api.HandlerPostSignup)
	router.GET("/logout", api.HandlerGetLogout)

	err := router.Run(
		helper.String(helper.Get(serverConfig, "host")) +
			":" +
			helper.String(helper.Get(serverConfig, "port")))
	if err != nil {
		log.Fatal(err)
	}
}
