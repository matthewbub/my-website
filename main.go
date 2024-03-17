package main

import (
	"matthewbub/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./apps/com/dist/index.html")
	r.Static("/assets", "./apps/com/dist/assets")

	r.GET("/", routes.ServeStatic)

	r.Run()
}
