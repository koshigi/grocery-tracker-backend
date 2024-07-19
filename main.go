package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	models "github.com/koshigi/grocery-tracker-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var (
	server       *gin.Engine
	DATABASE_URL string
	DB           *gorm.DB
)

func init() {
	// load environment variables
	env_err := godotenv.Load(".env")

	if env_err != nil {
		panic("Error loading .env file")
	}

	DATABASE_URL = os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		panic("DATABASE_URL environment variable is not set. Exiting.")
	}

	// set port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	os.Setenv("PORT", PORT)

	var db_err error
	DB, db_err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if db_err != nil {
		panic("Couldn't connect to database.")
	}

	server = gin.Default()
}

func main() {
	//Stone's initial test Get enpoint
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Placeholder for endpoints until someone figures out what is post, put, get

	// Creating the List parent object
	server.POST("/createList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "create list test",
		})
	})

	// Retrieving the list and its children
	server.GET("/getList/:id/", func(c *gin.Context) {
		id := c.Param("id")
		result := DB.First(&models.List{}, id)

		//TODO: fix this error Error #01: json: unsupported type: func() time.Time
		// unable to marshal the time data type, maybe specify in models?
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	// Deleting a list
	server.DELETE("/deleteList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete list test",
		})
	})

	// Update items to list
	server.POST("/updateItems", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "update item test",
		})
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	server.Run()
}
