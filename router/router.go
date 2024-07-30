package router

import (
	"blog-f1sh/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)

		v1.POST("/posts", controllers.CreatePost)
		v1.GET("/posts", controllers.GetPosts)
		v1.GET("/posts/:id", controllers.GetPost)

		v1.POST("/comments", controllers.CreateComment)
		v1.GET("/comments", controllers.GetComments)
	}

	return r
}
