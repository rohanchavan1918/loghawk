package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoutes(db *gorm.DB) *gin.Engine {

	// Default router
	router := gin.Default()
	router.Use(cors.Default())

	GetTagRoutes(db, router)
	GetTagRuleRoutes(db, router)
	GetIngestRoutes(db, router)
	GetLogRoutes(db, router)

	// Final router
	return router
}
