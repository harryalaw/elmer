package cmd

import (
	"fmt"
	"os"

	"github.com/harryalaw/elmer/internal/serialization"
)

type Query struct {
	Query string
}

func (q *Query) Exec() error {
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

	database, err := serialization.ImportDb(dbFilePath)
	if err != nil {
		return err
	}

	dir := database.Find(q.Query)

	if dir == nil {
		return fmt.Errorf("No such directory")
	}

	fmt.Println(dir.Path())
	return nil
}
