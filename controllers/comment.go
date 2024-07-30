package controllers

import (
	"blog-f1sh/config"
	"blog-f1sh/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func GetComments(c *gin.Context) {
	var comments []models.Comment
	if err := config.DB.Preload("User").Preload("Post").Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)

}
