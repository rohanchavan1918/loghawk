package routes

import (
	"loghawk/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLogRoutes(db *gorm.DB, router *gin.Engine) {
	router.GET("/logs", func(c *gin.Context) {
		var logs []models.Log
		// db.Find(&tags)
		db.Preload("Tag").Find(&logs)
		c.JSON(http.StatusOK, logs)
	})

	router.GET("/stats", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		rand.Seed(time.Now().UnixNano())

		// LogsParsed := rand.Intn(176)
		// errorCount := rand.Intn(12)
		// AlertsSent := errorCount - rand.Intn(3)

		LogsParsed := 1599
		errorCount := 56
		AlertsSent := 54

		c.JSON(http.StatusOK, gin.H{"logs_parsed": LogsParsed, "error_count": errorCount, "alerts_sent": AlertsSent})

	})

}
