package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import _ "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello Gin",
		})
	})

	r.Run(":9090")
}
