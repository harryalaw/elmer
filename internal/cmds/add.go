package cmd

import (
	"github.com/harryalaw/elmer/internal/db"
)

func AddCommand(pathName string, db *db.Db) *Command {
	return &Command{
		Name: "Add",
		Args: []string{pathName},
		Exec: add(pathName, db),
	}
}

func add(pathName string, database *db.Db) func() (*db.Db, error) {
	return func() (*db.Db, error) {
		dir := database.Find(pathName)
		if dir == nil {
			dir = db.NewDir(pathName)
			database = database.AddDir(dir)
		}

		database.Use(dir)

		return database, nil
	}
}
