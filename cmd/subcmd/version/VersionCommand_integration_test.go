// +build integration

package version_test

import (
	"fops/cmd"
	"fops/cmd/root"
	"fops/cmd/subcmd/version"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"strings"
)

func TestGetSubCommand(t *testing.T) {
	cmd, cmdErr := version.GetSubCommand()
	assert.NoError(t, cmdErr)
	assert.Equal(t, "*cobra.Command", reflect.TypeOf(cmd).String())
}

func TestChecksumRUNWithFile(t *testing.T) {
	setRootErr := root.SetRootCommand(&cobra.Command{
		Use: "fops",
		SilenceUsage: true,
		SilenceErrors: true,
	})
	rootCmd, getRootErr := root.GetRootCommand()
	command := "fops version"
	cmdArgs := strings.Split(command, " ")
	rootCmd.SetArgs(cmdArgs[1:])
	executeErr := cmd.Execute()
	assert.NoError(t, executeErr)
	assert.NoError(t, setRootErr)
	assert.NoError(t, getRootErr)
}