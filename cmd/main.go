package main

import (
	"fmt"
	"os"

	cmd "github.com/harryalaw/elmer/internal/cmds"
	"github.com/harryalaw/elmer/internal/db"
	"github.com/harryalaw/elmer/internal/serialization"
)

func main() {
	elmerPath := os.Getenv("ELMER_DB_PATH")

	if elmerPath == "" {
		panic("Please set a path for the elmer db, we use the 'ELMER_DB_PATH' value")
	}

	database, err := serialization.ImportDb(elmerPath)

	if err != nil {
		panic(fmt.Sprintf("Could not import DB: %+v\n", err))
	}

	cmd, err := parseArgs(os.Args, database)

	if err != nil {
		panic(fmt.Sprintf("Couldn't parse args correctly. Got error: %+v\n", err))
	}

	fmt.Printf("%+v\n", database)

	database, err = cmd.Exec()
	if err != nil {
		panic(fmt.Sprintf("Couldn't execute command correctly. Got error: %+v\n", err))
	}

	err = serialization.WriteDb(database, elmerPath)

	if err != nil {
		panic(fmt.Sprintf("Couldn't write database correctly. Got error: %+v\n", err))
	}
}

func parseArgs(args []string, database *db.Db) (*cmd.Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("No arguments provided")
	}

	switch args[1] {
	case "add":
		if len(args) != 3 {
			return nil, fmt.Errorf("Too many args provided for add")
		}
		return cmd.AddCommand(args[2], database), nil
	case "init":
		if len(args) != 3 {
			return nil, fmt.Errorf("Init needs the path for where the file gets made")
		}
		return cmd.Init(args[2]), nil
	}
	return nil, fmt.Errorf("No command found for args: %+v\n", args)
}
