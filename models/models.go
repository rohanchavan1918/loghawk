package models

import (
	"encoding/json"
	"time"
)

type Product struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
}

type Tag struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"name"`
	Description string    `json:"description"`
	Tag         string    `json:"tag" binding:"required"`
	SlackUrl    string    `json:"slack_url" binding:"required"`
	Rules       []TagRule `json:"tag_rules"`
	Logs        []Log     `gorm:"foreignKey:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TagRule struct {
	ID         uint   `json:"id" gorm:"primarykey"`
	TagId      int    `json:"tag_id"`
	Tag        Tag    `json:"tag" gorm:"foreignKey:TagId"`
	MatchType  string `json:"match_type"`
	MatchValue string `json:"match_value"`
	Priority   int    `json:"priority"`

	// Regex      string `json:"regex"`
	// StartsWith string `json:"starts_with"`
	// EndsWith   string `json:"ends_with"`
	// Contains   string `json:"contains"`
	// Custom     string `json:"drools"`
}

type CreateTagRuleRequest struct {
	TagID      int    `json:"tag_id" binding:"required"`
	MatchType  string `json:"match_type" binding:"required"`
	MatchValue string `json:"match_value" binding:"required"`
	Priority   int    `json:"priority"`
}

type Ingest struct {
	json.RawMessage
}
