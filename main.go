package main

import (
	"fmt"
	"log"
	"os"
	"store-inventory-management/src/config"
	"store-inventory-management/src/database"
	"store-inventory-management/src/server"
)

func main() {
	dsn := config.DatabaseUrlBuilder()
	db, err := database.New(dsn)
	defer database.Close(db)
	database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s := server.New(fmt.Sprintf(":%s", port), db)
	s.Serve()
}
