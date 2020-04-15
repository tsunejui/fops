package cmd

import (
	"fmt"
	"fops/cmd/root"
	"fops/cmd/subcmd/checksum"
	"fops/cmd/subcmd/linecount"
	"fops/cmd/subcmd/version"
	"fops/expection"
	"github.com/spf13/cobra"
)

var commandName string = "fops"

func init() {
	if err := root.SetRootCommand(&cobra.Command{Use: commandName}); err != nil {
		expection.ExitOsError(err)
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
		}else{
			cmd.AddCommand(subCmd)
		}
	}
	if executeErr := cmd.Execute(); executeErr != nil {
		return executeErr
	}
	return nil
}