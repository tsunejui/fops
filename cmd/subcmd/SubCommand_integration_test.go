// +build integration

package subcmd_test

import (
	"fops/cmd/subcmd"
	"fops/cmd/subcmd/version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubCommandGetShortName(t *testing.T) {
	var shortName string = "test"
	assert.Equal(t, (&subcmd.SubCommand{
		ShortName: shortName,
		Usage:     "",
	}).GetShortName(), shortName)
}
func TestSubCommandGetUsage(t *testing.T) {
	var usage string = "test"
	assert.Equal(t, (&subcmd.SubCommand{
		ShortName: "",
		Usage:     usage,
	}).GetUsage(), usage)
}

func TestGetCommandWithSubCommand(t *testing.T) {
	_, verCmdErr := version.GetSubCommand()
	assert.NoError(t, verCmdErr)
}