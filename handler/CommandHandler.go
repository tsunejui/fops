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

func (s *CommandHandler) GetCommands() []commands.CommandInterface {
	return s.Commands
}

func (s *CommandHandler) Do (args []string) error {
	var handle bool = false
	var cmdTemplae []templates.CommandTemplate
	if len(args) > 1 {
		for _, command := range s.GetCommands() {
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
			cmdTemplae = append(cmdTemplae, templates.CommandTemplate{
				Name:        command.GetCommandName(),
				Description: command.GetDescription(),
			})
		}
	}
	if ! handle {
		if rootCommandHandleErr := root.GetRootCommand().Do(args); rootCommandHandleErr != nil {
			//log.Println("Failed to parse root command: ", rootCommandHandleErr)
			templates.RootCommandUsage(cmdTemplae)
			root.GetRootCommand().Flags.PrintDefaults()
		}
	}
	return nil
}
