package version

import (
	"flag"
	"fops/commands"
	"github.com/stretchr/testify/assert"
	"reflect"
	"rsc.io/getopt"
	"testing"
)

func TestGetCommand(t *testing.T) {
	assert.Equal(t, "*version.command", reflect.TypeOf(GetCommand()).String())
}

func TestHandle(t *testing.T) {
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
	}).Handle([]string{}))
}

func TestInitial(t *testing.T) {
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
	}).Initial())
}