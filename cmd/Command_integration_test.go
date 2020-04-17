// +build integration

package cmd_test

import (
	"fops/cmd"
	"fops/cmd/root"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	setRootErr := root.SetRootCommand(&cobra.Command{
		Use: "fops",
		SilenceUsage: true,
		SilenceErrors: true,
	})
	rootCmd, getRootErr := root.GetRootCommand()
	command := "fops -h"
	cmdArgs := strings.Split(command, " ")
	rootCmd.SetArgs(cmdArgs[1:])
	executeErr := cmd.Execute()
	assert.NoError(t, executeErr)
	assert.NoError(t, setRootErr)
	assert.NoError(t, getRootErr)
}