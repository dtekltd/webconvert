package main

import (
	"log"
	"os"

	"github.com/tphan267/webconvert/pkg/converter"
)

func main() {
	// Create Fiber app
	app := converter.NewApp()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"
	}

	// Start server
	log.Fatal(app.Listen(port))
}
