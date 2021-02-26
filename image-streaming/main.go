package main

import (
	"github.com/gin-gonic/gin"
	"github.com/image-streaming/libs"
)

// https://stackoverflow.com/questions/22945486/golang-converting-image-image-to-byte

func stream(c *gin.Context) {
	img := libs.GetScreenshot()

	c.Data(200, "image/png", []byte(string(img)))
}

func main() {

	server := gin.Default()
	server.GET("/", stream)
	server.Run(":8080")
}
