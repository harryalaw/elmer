package cmd

import (
	"github.com/harryalaw/elmer/internal/db"
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
	// get the db from the file
	// add or update our thing

	// save the db
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
