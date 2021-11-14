package main

import (
	"fmt"
	// "net/http"

	"github.com/RohitKuwar/go_api_gin/routes"
)

func main() {
	fmt.Println("Server is runnig on port from config")
	r := routes.SetupRouter()
	r.Run(":8080")
}
