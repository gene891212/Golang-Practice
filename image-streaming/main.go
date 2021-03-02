package main

import (
	"github.com/gin-gonic/gin"
	"github.com/image-streaming/libs"
)

const header = "--frame\r\n" + "Content-Type: image/png\r\n\r\n"

func stream(c *gin.Context) {
	c.Header("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	for {
		img, err := libs.GetScreenshot()
		if err != nil {
			c.String(400, err.Error())
		}
		data := header + img.String()
		c.Writer.Write([]byte(data))
	}
}

// func stream(c *gin.Context) {
// 	c.Header("Content-Type", "multipart/x-mixed-replace; boundary=frame")
// 	chanStream := make(chan *bytes.Buffer, 10)
// 	go func() {
// 		defer close(chanStream)
// 		for {
// 			img, err := libs.GetScreenshot()
// 			if err != nil {
// 				c.String(400, err.Error())
// 			}
// 			chanStream <- img
// 		}
// 	}()

// 	c.Stream(func(w io.Writer) bool {
// 		if msg, ok := <-chanStream; ok {
// 			w.Write([]byte(header))
// 			w.Write(msg.Bytes())
// 			return true
// 		}
// 		return false
// 	})
// }

func main() {
	server := gin.Default()
	server.GET("/", stream)
	server.Run(":8000")
}
