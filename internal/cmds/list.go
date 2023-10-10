package cmd

import (
	"fmt"

	"github.com/harryalaw/elmer/internal/config"
)

type List struct {
}

func ListCommand() *List {
	return &List{}
}

func (l *List) Exec() error {
	db, err := config.LoadDb()

	if err != nil {
		return err
	}

	dirs := db.Dirs()

	for _, dir := range dirs {
		fmt.Printf("%s:  %d\n", dir.Path(), dir.Rank())
	}

	return nil
}
