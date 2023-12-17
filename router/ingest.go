package routes

import (
	"encoding/json"
	"fmt"
	"loghawk/parser"
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

		var data map[string]string
		err := json.Unmarshal(input, &data)
		if err != nil {
			fmt.Println("Err > ", err)
		}

		fmt.Println("Data > ", data)
		fmt.Println("Recieved logs from tag : ", data["tag"], data["log"])

		log, dataOk := data["log"]
		if dataOk {
			parser.ParseLogs(data["tag"], log)
		}

		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
}
