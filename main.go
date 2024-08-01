package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

func main() {
    // Create a Gin router
    r := gin.Default()

    // Define a route for the "Hello World" API
    fmt.Println("running...")

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Home",
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

    // Start the server on port 8080
    r.Run(":8080")
}

