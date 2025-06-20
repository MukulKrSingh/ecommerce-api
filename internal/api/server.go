package api

/* Purpose: Can contain the main application struct that wires up all the handlers, services, and repositories.
This centralizes the dependency injection and initialization of your application components.*/

import (
	"log"

	"github.com/ecommerce-api/config"
	"github.com/ecommerce-api/internal/api/rest"
	"github.com/ecommerce-api/internal/api/rest/handlers"
	"github.com/ecommerce-api/internal/helpers"
	"github.com/ecommerce-api/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupServer(cfg config.AppConfig) {

	//Initialize the router
	app := gin.Default()

	//Intialize the database connection if needed
	db, err := gorm.Open(postgres.Open("host=127.0.0.1 user=root password=root dbname=e-commerce port=5432 sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	log.Println("Database connection established successfully")

	//Run the database migrations

	e := db.AutoMigrate(&model.User{}, &model.Product{})
	if e != nil {
		log.Fatalf("Failed to run migrations: %v", e)
	}
	log.Println("Database migrations completed successfully")

	// Set up routes and Handlers
	// Example: router.GET("/products", productHandler.GetProducts)
	// You can add your handlers here, e.g.:
	// router.GET("/products", productHandler.GetProducts)
	// router.POST("/products", productHandler.CreateProduct)
	// router.PUT("/products/:id", productHandler.UpdateProduct)
	// router.DELETE("/products/:id", productHandler.DeleteProduct)

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the E-commerce API - HEALTH CHECK OK",
		})
	})

	auth := helpers.NewAuth(cfg.AppSecret)

	restHandler := &rest.RestHandler{
		App:  app,
		DB:   db, // Uncomment and set your database connection if needed
		Auth: auth,
	}
	// Initialize the REST API routes
	setUpRoutes(restHandler)

	app.Run(cfg.ServerPort)

}

func setUpRoutes(handler *rest.RestHandler) {

	handlers.InitUserRoutes(handler)
}
