package cmd

import (
	"fmt"
	"os/exec"

	"github.com/harryalaw/elmer/internal/config"
)

type Query struct {
	Query string
}

func (q *Query) Exec() error {
	database, err := config.LoadDb()
	if err != nil {
		return err
	}

	dir := database.Find(q.Query)

	if dir == nil {
		return fmt.Errorf("No such directory")
	}
	fmt.Println(dir.Path())

	cd := exec.Command("cd", dir.Path())
	cd.Run()
	return nil
}
