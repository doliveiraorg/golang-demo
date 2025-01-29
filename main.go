package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("templates/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Página Inicial",
        })
    })

    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    r.GET("/foo", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "bar",
        })
    })

    fmt.Println("Running...")

    r.Run(":8080")
}

