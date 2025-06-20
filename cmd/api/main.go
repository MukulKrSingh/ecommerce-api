package main

import (
	"github.com/ecommerce-api/config"
	"github.com/ecommerce-api/internal/api"
)

func main() {
	// This is a placeholder for the main function.
	// You can add your code here to run the application.

	appConfig, errr := config.SetupEnv()

	if errr != nil {
		panic("Error loading environment variables: " + errr.Error())
	}

	//Start the server
	api.SetupServer(appConfig)

}
