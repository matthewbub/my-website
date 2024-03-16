package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set the location of the templates
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")

	// Define a route to serve the HTML file
	r.GET("/", func(c *gin.Context) {
		// Render the HTML file
		c.HTML(200, "index.html", gin.H{
			"title": "Home Page",
		})
	})

	// TODO this should only serve if local
	// because yall dont need to be up in this
	r.GET("/admin", func(c *gin.Context) {
		// Render the HTML file
		c.HTML(200, "local--create-post.html", gin.H{
			"title": "Admin Page",
		})
	})

	// Handle the form submission
	r.POST("/submit-post", func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")

		// Here you would typically handle the data, e.g., validate and store it in the database
		log.Printf("Received new post with title: %s and content: %s", title, content)

		// Redirect or inform the user of success
		c.HTML(http.StatusOK, "local--create-post.html", gin.H{
			"title": "Post Submitted",
		})
	})

	r.Run()
}
