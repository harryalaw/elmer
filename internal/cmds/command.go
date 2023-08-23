package cmd

type Command interface {
	Exec() error
}
