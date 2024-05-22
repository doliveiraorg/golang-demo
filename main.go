package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create a Gin router
    r := gin.Default()

    // Define a route for the "Hello World" API
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // Start the server on port 8080
    r.Run(":8080")
}

