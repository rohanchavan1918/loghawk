package routes

import (
	"fmt"
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
		var count int64
		// LogsParsed := rand.Intn(176)
		// errorCount := rand.Intn(12)
		// AlertsSent := errorCount - rand.Intn(3)

		// LogsParsed := 1599
		// errorCount := 56
		// AlertsSent := 54
		// var count int64
		today := time.Now().Truncate(24 * time.Hour)
		db.Model(&models.Log{}).Where("created_at >= ?", today).Count(&count)
		fmt.Println("errorCount := ", count)
		fmt.Printf(" hit counts > %+v ", hitCounts)

		totalCount := 0
		for _, v := range hitCounts {
			totalCount += v
		}
		c.JSON(http.StatusOK, gin.H{"logs_parsed": totalCount, "error_count": count, "alerts_sent": count})
	})

}
