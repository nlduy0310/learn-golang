package controllers

import (
	"example/gin-gorm-crud/initializers"
	"example/gin-gorm-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func CreatePost(c *gin.Context) {
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	res := initializers.DB.First(&post, id)

	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	res := initializers.DB.First(&post, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	var content struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := c.Bind(&content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters",
		})
		return
	}

	post.Title = content.Title
	post.Body = content.Body

	res = initializers.DB.Save(&post)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while updating the post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated",
		"post":    post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	res := initializers.DB.First(&post, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	res = initializers.DB.Delete(&post)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while deleting post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
