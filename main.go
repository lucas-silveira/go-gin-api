package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
