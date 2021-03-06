package cmd

import (
	"errors"
	"fmt"
	"fops/cmd/root"
	"fops/cmd/subcmd/checksum"
	"fops/cmd/subcmd/linecount"
	"fops/cmd/subcmd/version"
	"fops/pkg/exception"

	"github.com/spf13/cobra"
)

var commandName string = "fops"

var executeCommand = func(cmd *cobra.Command) error {
	return cmd.Execute()
}

func init() {
	if err := root.SetRootCommand(&cobra.Command{
		Use: commandName,
		SilenceUsage: true,
		SilenceErrors: true,
	}); err != nil {
		exception.ExitOsError(err)
	}
}

func Execute() error {
	cmd, cmdErr := root.GetRootCommand()
	if cmdErr != nil {
		return cmdErr
	}
	for _, getCmd := range []func() (*cobra.Command, error){
		linecount.GetSubCommand,
		checksum.GetSubCommand,
		version.GetSubCommand,
	} {
		if subCmd, subCmdErr := getCmd(); subCmdErr != nil {
			fmt.Println(subCmdErr)
			continue
		} else {
			cmd.AddCommand(subCmd)
		}
	}
	if executeErr := executeCommand(cmd); executeErr != nil {
		fmt.Println(executeErr)
		return errors.New("error execute")
	}
	return nil
}
