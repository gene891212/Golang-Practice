package libs

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iot/stru"
)

// CreateUser ...
func CreateUser(c *gin.Context) {
	var user stru.UserInfo
	err := c.ShouldBindQuery(&user)
	if err == nil {
		InsertData(user)
		now, _ := time.Now().MarshalText()
		c.JSON(http.StatusOK, gin.H{
			"timestamp": string(now),
			"status":    http.StatusOK,
			"message":   "account create successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// message := stru.CreateSuccess{
	// 	Timestamp: string(now),
	// 	Status:    200,
	// 	Message:   "account created successful",
	// }
}
