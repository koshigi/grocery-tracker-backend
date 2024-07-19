package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/koshigi/grocery-tracker-backend/controllers"
	"github.com/koshigi/grocery-tracker-backend/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var (
	server       *gin.Engine
	DATABASE_URL string
	DB           *gorm.DB

	list_controller       controllers.ListController
	list_route_controller routers.ListRouteController
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

	// initalize controller and routecontrollers here
	list_controller = controllers.NewListController(DB)
	list_route_controller = routers.NewListRouteController(list_controller)

	server = gin.Default()
}

func main() {
	// main router group
	router := server.Group("/api")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	list_route_controller.ListRoute(router)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	server.Run()
}
