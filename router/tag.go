package routes

import (
	"fmt"
	"loghawk/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTagRoutes(db *gorm.DB, router *gin.Engine) {

	router.POST("/tags", func(c *gin.Context) {
		var tag models.Tag
		if err := c.ShouldBindJSON(&tag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&tag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tag)
	})

	// Read All (GET)
	router.GET("/tags", func(c *gin.Context) {
		var tags []models.Tag
		// db.Find(&tags)
		db.Preload("Rules").Find(&tags)
		c.JSON(http.StatusOK, tags)
	})

	// router.GET("/tags/rules", func(c *gin.Context) {
	// 	var allTagRules []models.

	// 	// Fetch all tags and their associated rules from the database
	// 	db.Preload("Rules").Find(&allTagRules)

	// 	c.JSON(http.StatusOK, allTagRules)
	// })

	// Read One (GET)
	router.GET("/tags/:id", func(c *gin.Context) {
		id := c.Param("id")
		var tag models.Tag
		db.First(&tag, "id = ?", id).Preload("Rules")

		if tag.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
			return
		}

		c.JSON(http.StatusOK, tag)
	})

	// Update (PUT)
	router.PUT("/tags/:id", func(c *gin.Context) {
		id := c.Param("id")
		_id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to delete logs.")
		}
		var updatedtag models.Tag
		if err := c.ShouldBindJSON(&updatedtag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Model(&models.Tag{ID: uint(_id)}).Updates(&updatedtag)

		c.JSON(http.StatusOK, updatedtag)
	})

	// Delete (DELETE)
	router.DELETE("/tags/:id", func(c *gin.Context) {
		id := c.Param("id")
		_id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to delete logs.")
		}
		db.Delete(&models.Tag{ID: uint(_id)})

		c.JSON(http.StatusOK, gin.H{"message": "tag deleted"})
	})
}

func GetTagRuleRoutes(db *gorm.DB, router *gin.Engine) {

	router.POST("/tag-rules", func(c *gin.Context) {

		var TagRuleReq models.CreateTagRuleRequest
		var tag models.Tag
		if err := c.ShouldBindJSON(&TagRuleReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Tag rule req > ", TagRuleReq)
		id := TagRuleReq.TagID
		result := db.First(&tag, id)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Tag does not exist"})
			return
		}

		fmt.Println("Tag > ", tag)
		// Create new TagRule
		TagRule := models.TagRule{
			TagId:      TagRuleReq.TagID,
			Tag:        tag,
			MatchType:  TagRuleReq.MatchType,
			MatchValue: TagRuleReq.MatchValue,
			Priority:   TagRuleReq.Priority,
		}

		if err := db.Create(&TagRule).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, TagRule)
	})

	// Read All (GET)
	router.GET("/tag-rules", func(c *gin.Context) {
		var tagRule []models.TagRule
		// db.Find(&tagRule)
		db.Preload("Tag").Find(&tagRule)
		c.JSON(http.StatusOK, tagRule)
	})

	// Read One (GET)
	router.GET("/tag-rules/:id", func(c *gin.Context) {
		id := c.Param("id")
		var tagRule models.TagRule
		db.First(&tagRule, "id = ?", id)

		if tagRule.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
			return
		}

		c.JSON(http.StatusOK, tagRule)
	})

	// Update (PUT)
	router.PUT("/tag-rules/:id", func(c *gin.Context) {
		id := c.Param("id")
		_id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to delete logs.")
		}
		var updatedtagRule models.TagRule
		if err := c.ShouldBindJSON(&updatedtagRule); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Model(&models.TagRule{ID: uint(_id)}).Updates(&updatedtagRule)

		c.JSON(http.StatusOK, updatedtagRule)
	})

	// Delete (DELETE)
	router.DELETE("/tagRule-rules/:id", func(c *gin.Context) {
		id := c.Param("id")
		_id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to delete logs.")
		}
		db.Delete(&models.TagRule{ID: uint(_id)})

		c.JSON(http.StatusOK, gin.H{"message": "tagRule deleted"})
	})
}
