// +build integration

package checksum_test

import (
	"fmt"
	"fops/cmd"
	"fops/cmd/root"
	"fops/cmd/subcmd/checksum"
	"fops/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/spf13/cobra"
	"reflect"
	"testing"
	"strings"
)

func TestGetSubCommand(t *testing.T) {
	cmd, cmdErr := checksum.GetSubCommand()
	assert.NoError(t, cmdErr)
	assert.Equal(t, "*cobra.Command", reflect.TypeOf(cmd).String())
}

func TestChecksumRUWithFile(t *testing.T) {
	initSystemErr := service.InitialSystemService()
	setRootErr := root.SetRootCommand(&cobra.Command{
		Use: "fops",
		SilenceUsage: true,
		SilenceErrors: true,
	})
	rootCmd, getRootErr := root.GetRootCommand()
	filePath := "../../../test/myfile.txt"
	command := fmt.Sprintf("fops checksum -f %s --md5", filePath)
	cmdArgs := strings.Split(command, " ")
	rootCmd.SetArgs(cmdArgs[1:])
	executeErr := cmd.Execute()
	assert.NoError(t, executeErr)
	assert.NoError(t, initSystemErr)
	assert.NoError(t, setRootErr)
	assert.NoError(t, getRootErr)
}