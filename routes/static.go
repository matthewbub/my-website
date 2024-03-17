package routes

import "github.com/gin-gonic/gin"

func ServeStatic(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
