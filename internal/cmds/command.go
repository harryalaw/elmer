package cmd

import "github.com/harryalaw/elmer/internal/db"

type Command struct {
	Name string
	Args []string
	Exec func() (*db.Db, error)
}
