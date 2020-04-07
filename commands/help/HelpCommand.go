package help

import (
	"flag"
	"fmt"
	"fops/commands"
	"fops/handler"
	"fops/templates"
	"rsc.io/getopt"
)

type command struct {
	commands.SubCommand
}
func (c *command) Initial () {

}
func (c *command) Handle (args []string) error {
	if len(args) > 2 {
		for _, command := range handler.GetCommandHandler().Commands {
			if args[2] == command.GetCommandName() {
				fmt.Println(command.GetDescription())
				command.Initial()
				templates.SubCommandUsage(command.GetDescription(), command.GetCommandName())
				command.GetField().PrintDefaults()
				break
			}
		}
	}else{
		templates.HelpCommandUsage(c.GetDescription())
	}
	return nil
}
func GetCommand() commands.CommandInterface {
	help := getopt.NewFlagSet("help", flag.ExitOnError)
	return &command{
		commands.SubCommand{
			Flag: help,
			CommandName: "help",
			Description: "Help about commands",
		},
	}
}