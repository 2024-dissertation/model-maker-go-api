package seeds_test

import (
	"os"
	"testing"

	"github.com/Soup666/diss-api/seeds"
	"github.com/stretchr/testify/assert"
)

func setupTestDirs(t *testing.T) {
	os.MkdirAll("./uploads", os.ModePerm)
	os.MkdirAll("./objects", os.ModePerm)
	os.WriteFile("./uploads/test.txt", []byte("upload"), 0644)
	os.WriteFile("./objects/test.txt", []byte("object"), 0644)
}

func teardownTestDirs() {
	os.RemoveAll("./uploads")
	os.RemoveAll("./objects")
}

func TestMakeBackup(t *testing.T) {
	setupTestDirs(t)
	defer teardownTestDirs()

	err := seeds.MakeBackup()
	assert.NoError(t, err)

	// Check backup directory contains the moved folders
	_, errUploads := os.Stat("./backup/uploads")
	_, errObjects := os.Stat("./backup/objects")

	assert.NoError(t, errUploads)
	assert.NoError(t, errObjects)
}

func TestCopyFiles(t *testing.T) {
	// Setup dummy tar.gz
	os.MkdirAll("./seeds/backup", os.ModePerm)

	defer func() {
		os.Remove("./backup.tar.gz")
	}()

	err := seeds.CopyFiles()
	assert.NoError(t, err)

	_, err = os.Stat("./backup.tar.gz")
	assert.True(t, os.IsNotExist(err)) // should be removed after extraction step
}

func TestAllSeeds(t *testing.T) {
	seedsList := seeds.All()
	assert.Len(t, seedsList, 3)
	assert.Equal(t, "CreateTestUser", seedsList[0].Name)
	assert.NotNil(t, seedsList[0].Run)
}
