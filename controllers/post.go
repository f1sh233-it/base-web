package controllers

import (
	"blog-f1sh/config"
	"blog-f1sh/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Preload("User").Find(&posts).Error; err != nil {
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	var post models.Post
	if err := config.DB.Preload("User").First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
