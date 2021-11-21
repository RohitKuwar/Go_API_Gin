package main

import (
	"fmt"
	"log"

	"github.com/RohitKuwar/go_api_gin/config"
	"github.com/RohitKuwar/go_api_gin/routes"
	"github.com/joho/godotenv"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println("Server is successfully runnig on port:", config.Port)
	log.Printf("Server is successfully runnig on port:", config.Port)
	r := routes.SetupRouter()
	r.Run(":" + config.Port)
}
