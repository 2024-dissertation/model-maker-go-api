package main

import (
	"log"
	"os"

	"github.com/Soup666/diss-api/seeds"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Starting seeding database...")
	DB, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	for _, seed := range seeds.All() {
		log.Printf("Running seed '%s'", seed.Name)
		if err := seed.Run(DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	log.Println("Seeding completed successfully")

	log.Println("Backing up files...")
	if err := seeds.MakeBackup(); err != nil {
		log.Fatalf("Error backing up files: %s", err)
	}

	log.Println("Moving files...")
	if err := seeds.CopyFiles(); err != nil {
		log.Fatalf("Error moving files: %s", err)
	}

	log.Println("Files moved successfully")
	log.Println("Seeding and file operations completed successfully")
}
