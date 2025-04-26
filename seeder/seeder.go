package main

import (
	"log"

	"github.com/Soup666/diss-api/database"
	"github.com/Soup666/diss-api/seeds"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	for _, seed := range seeds.All() {
		log.Printf("Running seed '%s'", seed.Name)
		if err := seed.Run(database.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	log.Println("Seeding completed successfully")
}
