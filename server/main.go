package main

import (
	"os"
)

func main() {
	addr := os.Getenv("CACHE_ADDR")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASSWORD")
	if addr == "" {
		addr = "localhost:3124"
	}

	sv := NewServer(db_host, db_port, db_password)

	sv.Run("3000")
}
