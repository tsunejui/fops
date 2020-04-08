package commands

import (
	"rsc.io/getopt"
)

type CommandInterface interface {
	GetField() *getopt.FlagSet
	GetDescription() string
	GetCommandName() string
	Parse([]string) error
	Initial() error
	Handle([]string) error
}

type SubCommand struct {
	Flag *getopt.FlagSet
	Description string
	CommandName string
}

func (s *SubCommand) GetField() *getopt.FlagSet {
	return s.Flag
}

func (s *SubCommand) GetDescription() string {
	return s.Description
}

func (s *SubCommand) GetCommandName() string {
	return s.CommandName
}

func (s *SubCommand) Parse(args []string) error {
	return s.GetField().Parse(args)
}