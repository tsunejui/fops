package subcmd

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"reflect"
	"rsc.io/getopt"
	"testing"
)

func TestGetCommandName(t *testing.T) {
	assert.Equal(t, "string", reflect.TypeOf((&SubCommand{
		Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
		Description: "",
		CommandName: "",
	}).GetCommandName()).String())
}

func TestGetDescription(t *testing.T) {
	assert.Equal(t, "string", reflect.TypeOf((&SubCommand{
		Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
		Description: "",
		CommandName: "",
	}).GetDescription()).String())
}

func TestGetField(t *testing.T) {
	assert.Equal(t, "*getopt.FlagSet", reflect.TypeOf((&SubCommand{
		Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
		Description: "",
		CommandName: "",
	}).GetField()).String())
}

func TestParse(t *testing.T) {
	assert.NoError(t, (&SubCommand{
		Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
		Description: "",
		CommandName: "",
	}).Parse([]string{}))
}

