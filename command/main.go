package main

import (
	"gogolook/router"
	"log"
	"os"
)

func main() {
	server := router.SetupRouter()

	port := ":8000"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	log.Println("port = ", port)
	server.Run(port)
}
