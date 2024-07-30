package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	PostID  uint   `gorm:"primaryKey" json:"post_id"`
	UserID  uint   `gorm:"primaryKey" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Post    Post   `gorm:"foreignKey:PostID" json:"post"`
}
