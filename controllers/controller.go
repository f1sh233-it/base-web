package controllers

import (
	"blog-f1sh/config"
	"blog-f1sh/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNewUser(c *gin.Context) {
	// 获取当前登录的用户
	currentUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}

	// 确保 currentUser 类型正确
	user, ok := currentUser.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User type assertion failed"})
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置 post 的 UserID
	post.UserID = user.ID

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}
