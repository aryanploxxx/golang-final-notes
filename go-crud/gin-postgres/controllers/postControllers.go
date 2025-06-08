package controllers

import (
	"gin-postgres/initializers"
	"gin-postgres/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct { // could have been another name
		Body  string
		Title string
	}

	c.Bind(&body) // this command will bind the incoming request body to the body variable
	// In this code, the body variable is populated with the request body data by calling the c.Bind(&body) method from the Gin framework.

	// newPost := models.Post{Title: "Hello", Body: "World"}
	// hardcoded method to create a new post
	newPost := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&newPost)
	// directly convert the struct to a row in table

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"created new post": newPost})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts) // find all the posts and initialize the posts variable with the result
	c.JSON(200, gin.H{"all posts": posts})
}

func GetPostsByID(c *gin.Context) {
	// get the id from the URL
	id := c.Param("id")

	// get the post
	var post models.Post
	initializers.DB.Find(&post, id) // find all the posts and initialize the posts variable with the result

	// return the post
	c.JSON(200, gin.H{"fetched post": post})
}

func PostUpdateByID(c *gin.Context) {
	// get the id from the URL
	id := c.Param("id")

	// get the data from the request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body) // this command will bind the incoming request body to the body variable

	// find the post we want to update
	var post models.Post
	initializers.DB.Find(&post, id) // find all the posts and initialize the posts variable with the result

	// update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	// post.Title = body.Title
	// post.Body = body.Body

	// // save the post
	// initializers.DB.Save(&post)

	// return the post
	c.JSON(200, gin.H{"updated post": post})
}

func PostDeleteByID(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
