package main

import (
	"net/http"

	"github.com/iot/libs"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}
func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.Static("/static", "./static")
	server.GET("/", test)
	server.POST("/create", libs.CreateUser)
	server.Run(":12345")
}
