package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//index page
	r.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "try to /flag",
		})
	})

	// authentication setup
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	// flag page
	authorized.GET("/flag", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"secret": "U0VDUkVUe3RoMTVfMTVfczNjcjN0X20zNTUzZzN9",
		})
	})

	r.Run(":8081")
}
