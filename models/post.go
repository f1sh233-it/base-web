package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"size:255" json:"title"`
	Content string `gorm:"size:255" json:"content"`
	UserID  uint   `gorm:"index" json:"user_id"`
	User    User   `gorm:"foreigner:UserID" json:"user"`
}
