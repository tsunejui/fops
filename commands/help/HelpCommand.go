package help

import (
	"flag"
	"fops/commands"
	"fops/handler"
	"fops/templates"
	"log"
	"rsc.io/getopt"
)

type command struct {
	commands.SubCommand
}
func (c *command) Initial () error {
	return nil
}
func (c *command) Handle (args []string) error {
	if len(args) > 2 {
		for _, command := range handler.GetCommandHandler().GetCommands() {
			if args[2] == command.GetCommandName() {
				if initialErr := command.Initial(); initialErr != nil {
					log.Println("Failed to intiial command: ", initialErr)
					break
				}
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