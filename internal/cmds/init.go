package cmd

import (
	"fmt"
	"os"

	"github.com/harryalaw/elmer/internal/db"
	"github.com/harryalaw/elmer/internal/serialization"
)

type Init struct {
}

// uses this if the ELMER_DIR_PATH is not set
const DEFAULT_DB_DIR = ".elmer"
const ELMER_FILE = "elmer.db"

func (i *Init) Exec() error {
	// check that db_dir exists else create it
	dbDir := os.Getenv("ELMER_DIR_PATH")

	if dbDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting user's home directory: %v\n", err)
			return err
		}
		dbDir = fmt.Sprintf("%s/%s", homeDir, DEFAULT_DB_DIR)
	}

	dbFilePath := fmt.Sprintf("%s/%s", dbDir, ELMER_FILE)

	// try to read the elmer.db file. If it is valid don't kill
	_, err := serialization.ImportDb(dbFilePath)

	// assuming that if this fails then the current DB is invalid
	// and we should make a new one
	if err == nil {
		return nil
	}

	database := db.FromDirs([]db.Dir{})
	return serialization.WriteDb(database, dbFilePath)
}
