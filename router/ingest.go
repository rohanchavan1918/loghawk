package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetIngestRoutes(db *gorm.DB, router *gin.Engine) {

	router.POST("/ingest", func(c *gin.Context) {
		var input json.RawMessage
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data map[string]interface{}
		err := json.Unmarshal(input, &data)
		if err != nil {
			fmt.Println("Err > ", err)
		}

		fmt.Println("Recieved logs from tag : ", data["tag"])

		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
}
