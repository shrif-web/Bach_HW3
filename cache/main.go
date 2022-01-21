package main

import (
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3124"
	}

	server := NewCacheServer(20)
	server.Start(port)
}
