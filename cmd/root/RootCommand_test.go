package root

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"reflect"
	"rsc.io/getopt"
	"testing"
)

func TestGetCommand(t *testing.T) {
	assert.Equal(t, "*root.RootCommand", reflect.TypeOf(GetRootCommand()).String())
}

func TestInitialRootCommand(t *testing.T) {
	assert.NoError(t, InitialRootCommand())
}

func TestDo(t *testing.T) {
	assert.Error(t, (&RootCommand{
		Flags: getopt.NewFlagSet("test", flag.ExitOnError),
	}).Do([]string{"test", "test"}))
}