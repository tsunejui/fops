package root

import (
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func SetRootCommand(cmd *cobra.Command) error {
	rootCmd = cmd
	return nil
}

func GetRootCommand() (*cobra.Command, error) {
	return rootCmd, nil
}
