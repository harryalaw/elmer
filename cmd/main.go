package main

import (
	"fmt"
	"os"

	cmd "github.com/harryalaw/elmer/internal/cmds"
)

func main() {
	cmd, err := parseArgs(os.Args)

	if err != nil {
		panic(fmt.Sprintf("Couldn't parse args correctly. Got error: %+v\n", err))
	}

	err = cmd.Exec()
	if err != nil {
		panic(fmt.Sprintf("Couldn't execute command correctly. Got error: %+v\n", err))
	}
}

func parseArgs(args []string) (cmd.Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("No arguments provided")
	}

	switch args[1] {
	case "add":
		if len(args) != 3 {
			return nil, fmt.Errorf("Too many args provided for add")
		}
		return cmd.AddCommand(args[2]), nil
	case "init":
		if len(args) != 2 {
			return nil, fmt.Errorf("init takes no args")
		}
		return &cmd.Init{}, nil
	case "query":
		if len(args) != 3 {
			return nil, fmt.Errorf("query needs a value to search for")
		}
		return &cmd.Query{Query: args[2]}, nil

	}
	return nil, fmt.Errorf("No command found for args: %+v\n", args)
}
