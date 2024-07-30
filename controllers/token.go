package controllers

import (
	"blog-f1sh/config"
	"blog-f1sh/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Token(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 这里需要确保 UserId 字段在 models.Post 中存在
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}
