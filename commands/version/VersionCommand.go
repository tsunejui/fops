package version

import (
	"flag"
	"fmt"
	"fops/commands"
	"rsc.io/getopt"
)

type command struct {
	commands.SubCommand
}

func (c *command) Initial () {

}

func (c *command) Handle (args []string) error {
	fmt.Println("fops v0.0.1")
	return nil
}

func GetCommand() commands.CommandInterface {
	version := getopt.NewFlagSet("version", flag.ExitOnError)
	return &command{
		commands.SubCommand{
			Flag: version,
			CommandName: "version",
			Description: "Show the version info",
		},
	}
}
