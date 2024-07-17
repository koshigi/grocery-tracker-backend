package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Documentation for the Gin library used to set up the endpoints
// https://github.com/gin-gonic/gin/blob/master/docs/doc.md
func main() {
	//Creation of a gin router with default middlewware
	r := gin.Default()

	//Stone's initial test Get enpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Placeholder for endpoints until someone figures out what is post, put, get

	// Creating the List parent object
	r.POST("/createList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "create list test",
		})
	})

	// Retrieving the list and its children
	r.GET("/getList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get list test",
		})
	})

	// Deleting a list
	r.DELETE("/deleteList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete list test",
		})
	})

	// Update items to list
	r.POST("/updateItems", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "update item test",
		})
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
