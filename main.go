package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	models "github.com/koshigi/grocery-tracker-backend/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		panic("DATABASE_URL environment variable is not set. Exiting.")
	}
}

func main() {
	//Creation of a gin router with default middlewware
	r := gin.Default()

	DATABASE_URL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database.")
	}

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
	r.GET("/getList/:id/", func(c *gin.Context) {
		id := c.Param("id")
		result := db.First(&models.List{}, id)

		//TODO: fix this error Error #01: json: unsupported type: func() time.Time
		// unable to marshal the time data type, maybe specify in models?
		c.JSON(http.StatusOK, gin.H{
			"data": result,
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
