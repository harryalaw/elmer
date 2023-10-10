package cmd

import (
	"os"

	"github.com/harryalaw/elmer/internal/config"
	"github.com/harryalaw/elmer/internal/db"
)

type Add struct {
	path string
}

func AddCommand() *Add {
	return &Add{}
}

func (a *Add) Exec() error {
	db, err := config.LoadDb()

	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	db, err = add(wd, db)

	if err != nil {
		return err
	}

	err = config.SaveDb(db)
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
