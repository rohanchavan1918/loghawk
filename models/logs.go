package models

import "time"

type Log struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Message   string `json:"msg"`
	Tag       Tag    `gorm:"foreignKey:TagID"`
	TagID     uint   `json:"tag_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
