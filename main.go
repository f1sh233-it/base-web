package main

import (
	"blog-f1sh/config"
	"blog-f1sh/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	//	config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	r := router.SetupRouter()

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	//r.Static("/static", "./frontend")
	r.Run()
}
