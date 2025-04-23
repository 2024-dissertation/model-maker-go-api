package seeds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Soup666/diss-api/seed"
	"github.com/Soup666/diss-api/utils"
	"gorm.io/gorm"
)

func MakeBackup() error {
	srcDirs := [2]string{"./uploads", "./objects"}
	destDir := "./backup/"

	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, srcDir := range srcDirs {
		os.Rename(srcDir, filepath.Join(destDir, filepath.Base(srcDir)))
	}

	return nil
}

func CopyFiles() error {
	util := utils.NewFileUtil()

	// Example paths
	currentDir, _ := os.Getwd()
	srcFile := filepath.Join(currentDir, "seeds", "backup", "backup.tar.gz")
	destFile := filepath.Join(currentDir, "backup.tar.gz")

	// Step 1: Copy tar.gz
	if err := util.CopyFile(srcFile, destFile); err != nil {
		fmt.Printf("Copy failed: %v\n", err)
		return nil
	}
	fmt.Println("Tar.gz copied successfully.")

	// Step 2: Extract
	f, err := os.Open(destFile)
	if err != nil {
		fmt.Printf("Failed to open tar.gz: %v\n", err)
		return nil
	}
	defer f.Close()

	if err := util.ExtractTarGz(f, currentDir); err != nil {
		fmt.Printf("Extraction failed: %v\n", err)
		return nil
	}

	// Step 3: Remove tar.gz

	if err := os.Remove(destFile); err != nil {
		fmt.Printf("Failed to remove tar.gz: %v\n", err)
		return nil
	}

	fmt.Println("Extraction successful.")

	return nil
}

func All() []seed.Seed {
	return []seed.Seed{
		{
			Name: "CreateTestUser",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Seed User", "KQmrXe88TwebIMh6AkbEV251Aec2")
			},
		},
		{
			Name: "CreateTestTask",
			Run: func(db *gorm.DB) error {
				return CreateTask(db, "Seed Figure", "This is seeded data to represent an example scan", true, 1)
			},
		},
		{
			Name: "CreateTestFiles",
			Run: func(db *gorm.DB) error {
				return CreateDummyFiles(db)
			},
		},
	}
}
