package cmd

import (
	"fmt"

	"github.com/harryalaw/elmer/internal/db"
	"github.com/harryalaw/elmer/internal/serialization"
)

func Init(elmerFilePath string) *Command {
	return &Command{
		Name: "Init",
		Args: []string{elmerFilePath},
		Exec: func() func() (*db.Db, error) {
			return func() (*db.Db, error) {
				database := db.FromDirs([]db.Dir{})
				if elmerFilePath == "" {
					return nil, fmt.Errorf("No value set for ELMER_DB_PATH")
				}

				if err := serialization.WriteDb(database, elmerFilePath); err != nil {
					return nil, err
				}

                fmt.Printf("%+v\n", database);
				return database, nil
			}
		}(),
	}
}
