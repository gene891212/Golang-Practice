package main

import (
	"net/http"

	"github.com/iot/libs"
	"github.com/iot/stru"

	"github.com/gin-gonic/gin"
)

func find(c *gin.Context) {
	result := libs.FindAccount(c.Query("findAccount"))

	data := stru.IndexData{
		AllAccount: result,
	}

	c.HTML(http.StatusOK, "index.html", data)
}
func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.Static("/static", "./static")
	server.GET("/", find)
	server.POST("/create", libs.CreateUser)
	server.Run(":12345")
}
