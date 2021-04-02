package main

import (
	"gin-api-server/lib"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetTime(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":    200,
		"timestamp": time.Now().Unix(),
	})
}

func GetData(c *gin.Context) {
	allUser := lib.DataFromDB()
	c.JSON(200, gin.H{
		"status":  200,
		"allUser": allUser,
	})
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", GetTime)
		api.GET("/data", GetData)
	}
	r.Run(":8001")
}
