package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

        r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.LoadHTMLFiles("static/index.html")

	r.Run(":8080") 
}

