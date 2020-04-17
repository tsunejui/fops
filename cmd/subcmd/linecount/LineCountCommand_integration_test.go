// +build integration

package linecount_test

import (
	"fmt"
	"fops/cmd"
	"fops/cmd/root"
	"fops/cmd/subcmd/linecount"
	"fops/pkg/service"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestGetSubCommand(t *testing.T) {
	cmd, cmdErr := linecount.GetSubCommand()
	assert.NoError(t, cmdErr)
	assert.Equal(t, "*cobra.Command", reflect.TypeOf(cmd).String())
}

func TestChecksumRUNWithFile(t *testing.T) {
	initSystemErr := service.InitialSystemService()
	setRootErr := root.SetRootCommand(&cobra.Command{
		Use: "fops",
		SilenceUsage: true,
		SilenceErrors: true,
	})
	rootCmd, getRootErr := root.GetRootCommand()
	filePath := "../../../test/myfile.txt"
	command := fmt.Sprintf("fops linecount -f %s", filePath)
	cmdArgs := strings.Split(command, " ")
	rootCmd.SetArgs(cmdArgs[1:])
	executeErr := cmd.Execute()
	assert.NoError(t, executeErr)
	assert.NoError(t, initSystemErr)
	assert.NoError(t, setRootErr)
	assert.NoError(t, getRootErr)
}

func TestChecksumRUNWithBinaryFile(t *testing.T) {
	initSystemErr := service.InitialSystemService()
	setRootErr := root.SetRootCommand(&cobra.Command{
		Use: "fops",
		SilenceUsage: true,
		SilenceErrors: true,
	})
	rootCmd, getRootErr := root.GetRootCommand()
	filePath := "../../../test/fops"
	command := fmt.Sprintf("fops linecount -f %s", filePath)
	cmdArgs := strings.Split(command, " ")
	rootCmd.SetArgs(cmdArgs[1:])
	executeErr := cmd.Execute()
	assert.Error(t, executeErr)
	assert.Equal(t, "error execute",executeErr.Error())
	assert.NoError(t, initSystemErr)
	assert.NoError(t, setRootErr)
	assert.NoError(t, getRootErr)
}

func TestChecksumRUNWithInvalidFile(t *testing.T) {
	initSystemErr := service.InitialSystemService()
	setRootErr := root.SetRootCommand(&cobra.Command{
		Use: "fops",
		SilenceUsage: true,
		SilenceErrors: true,
	})
	rootCmd, getRootErr := root.GetRootCommand()
	filePath := "../../../test/test"
	command := fmt.Sprintf("fops linecount -f %s", filePath)
	cmdArgs := strings.Split(command, " ")
	rootCmd.SetArgs(cmdArgs[1:])
	executeErr := cmd.Execute()
	assert.Error(t, executeErr)
	assert.Equal(t, "error execute",executeErr.Error())
	assert.NoError(t, initSystemErr)
	assert.NoError(t, setRootErr)
	assert.NoError(t, getRootErr)
}