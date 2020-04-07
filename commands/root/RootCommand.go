package root

import (
	"errors"
	"flag"
	"rsc.io/getopt"
)
type RootCommand struct {
	Flags *getopt.FlagSet
}
var rootCommand *RootCommand
var help bool
func InitialRootCommand() error {
	flags := getopt.NewFlagSet("main", flag.ExitOnError)
	flags.BoolVar(&help, "h", false, "Help about commands")
	flags.Alias("h", "help")
	rootCommand = &RootCommand {
		Flags: flags,
	}
	return nil
}
func GetRootCommand () *RootCommand{
	return rootCommand
}
func (r RootCommand) Do(args []string) error {
	if parseErr := r.Flags.Parse(args[1:]); parseErr != nil {
		return errors.New("Failed to parse root command")
	}
	if help {
		return errors.New("Show command help")
	}
	return nil
}
