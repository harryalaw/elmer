package config

import (
	"fmt"
	"os"

	"github.com/harryalaw/elmer/internal/db"
	"github.com/harryalaw/elmer/internal/serialization"
)

const DEFAULT_DB_DIR = ".elmer"
const ELMER_FILE = "elmer.db"

func dataDir() (string, error) {
	dbDir := os.Getenv("_ELMER_DATA_DIR")

	if dbDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting user's home directory: %v\n", err)
			return "", err
		}
		return homeDir, nil
	}
	return dbDir, nil
}

func databaseFilePath() string {
	dbDir, err := dataDir()
	if err != nil {
		panic(fmt.Errorf("Could not find data dir: %+v", err))
	}

	return fmt.Sprintf("%s/%s", dbDir, ELMER_FILE)
}

func LoadDb() (*db.Db, error) {
	dbFilePath := databaseFilePath()

	db, err := serialization.ImportDb(dbFilePath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveDb(database *db.Db) error {
	dbFilePath := databaseFilePath()

	err := serialization.WriteDb(database, dbFilePath)
	return err
}
