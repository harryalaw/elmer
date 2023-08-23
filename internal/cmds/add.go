package cmd

import (
	"fmt"
	"os"

	"github.com/harryalaw/elmer/internal/db"
	"github.com/harryalaw/elmer/internal/serialization"
)

type Add struct {
	path string
}

func AddCommand(pathName string) *Add {
	return &Add{
		path: pathName,
	}
}

func (a *Add) Exec() error {
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

	db, err := serialization.ImportDb(dbFilePath)
	if err != nil {
		return err
	}

	db, err = add(a.path, db)

	if err != nil {
		return err
	}

	err = serialization.WriteDb(db, dbFilePath)
	if err != nil {
		return err
	}

	return nil
}

func add(pathName string, database *db.Db) (*db.Db, error) {
	dir := database.Find(pathName)
	if dir == nil {
		dir = db.NewDir(pathName)
		database = database.AddDir(dir)
	}

	database.Use(dir)

	return database, nil
}
