package handler

import (
	"fops/commands"
	"fops/commands/root"
	"fops/templates"
)

type CommandHandler struct {
	Commands []commands.CommandInterface
	RootCommand *root.RootCommand
}

var commandHandler *CommandHandler

func InitialCommandHandler() error {
	commandHandler = &CommandHandler{}
	return nil
}

func GetCommandHandler() *CommandHandler{
	return commandHandler
}

func (s *CommandHandler) SetRootCommand(command *root.RootCommand) *CommandHandler {
	s.RootCommand = command
	return s
}

func (s *CommandHandler) AddCommand(command commands.CommandInterface) *CommandHandler {
	s.Commands = append(s.Commands, command)
	return s
}

func (s *CommandHandler) Do (args []string) error {
	var handle bool = false
	if len(args) > 1 {
		for _, command := range s.Commands {
			if args[1] == command.GetCommandName() {
				command.Initial()
				if parseErr := command.Parse(args[2:]); parseErr != nil {
					continue
				}
				if handleErr := command.Handle(args); handleErr != nil {
					templates.SubCommandUsage(command.GetDescription(), command.GetCommandName())
					command.GetField().PrintDefaults()
					break
				}
				handle = true
				break
			}
		}
	}
	if ! handle {
		if rootCommandHandleErr := root.GetRootCommand().Do(args); rootCommandHandleErr != nil {
			//log.Println("Failed to parse root command: ", rootCommandHandleErr)
			var cmdTemplae []templates.CommandTemplate
			for _, command := range s.Commands {
				cmdTemplae = append(cmdTemplae, templates.CommandTemplate{
					Name:        command.GetCommandName(),
					Description: command.GetDescription(),
				})
			}
			templates.RootCommandUsage(cmdTemplae)
			root.GetRootCommand().Flags.PrintDefaults()
		}
	}
	return nil
}
