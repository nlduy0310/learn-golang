package main

import (
	"os"

	"example/gin-gorm-crud/controllers"
	"example/gin-gorm-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.GetAllPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run(os.Getenv("RUN_ADDR"))
}
