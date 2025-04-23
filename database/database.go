package database

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Soup666/diss-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(connectionString string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return err
	}

	return nil
}

func SetupTestDB(t *testing.T) error {

	log.Println("Connecting to database...")
	log.Println(os.Getenv("DATABASE_URL"))

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"))
	err := ConnectDatabase(dsn)

	ResetTestDB()

	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
		return err
	}

	createEnumCommand := `
	--create types
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'taskstatus') THEN
        CREATE TYPE TaskStatus AS ENUM
        (
            'SUCCESS', 'INPROGRESS', 'FAILED', 'INITIAL'
        );
    END IF;
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reporttype') THEN
        CREATE TYPE ReportType AS ENUM
        (
            'BUG', 'FEEDBACK'
        );
    END IF;
END$$;
`

	DB.Exec(createEnumCommand)

	err = DB.AutoMigrate(&model.User{}, &model.Task{}, &model.Report{}, &model.Collection{})

	if err != nil {
		t.Fatalf("Failed to migrate test DB: %v", err)
		return err
	}
	return nil
}

func ResetTestDB() {
	log.Println("Resetting test database...")

	tables := []string{"users", "tasks", "reports", "collections"}
	for _, table := range tables {
		truncateTableCommand := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)
		DB.Exec(truncateTableCommand)
	}
}
