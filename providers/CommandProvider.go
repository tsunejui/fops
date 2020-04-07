package providers

import (
	"fops/commands"
	"fops/commands/checksum"
	"fops/commands/help"
	"fops/commands/linecount"
	"fops/commands/version"
)

func GetVersionCommand () commands.CommandInterface {
	return version.GetCommand()
}

func GetLineCountCommand() commands.CommandInterface {
	return linecount.GetCommand()
}

func GetHelpCommand() commands.CommandInterface {
	return help.GetCommand()
}

func GetChecksumCommand() commands.CommandInterface {
	return checksum.GetCommand()
}