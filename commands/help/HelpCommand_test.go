package help

import (
	"flag"
	"fops/commands"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"rsc.io/getopt"
	"testing"
)

type mockSubCommand struct {
	mock.Mock
	commands.SubCommand
}

func (m *mockSubCommand) Initial() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockSubCommand) Handle(handleArgs []string) error {
	args := m.Called(handleArgs)
	return args.Error(0)
}

func TestGetCommand(t *testing.T) {
	assert.Equal(t, "*help.command", reflect.TypeOf(GetCommand()).String())
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

func TestHandle(t *testing.T) {
	commandName := "test"
	subCommand := &mockSubCommand{
		SubCommand: commands.SubCommand{
			Flag: getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: commandName,
		},
	}
	subCommand.On("Initial").Return(nil)
	defer gostub.StubFunc(&getAllCommands, []commands.CommandInterface{
		subCommand,
	}).Reset()
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
	}).Handle([]string{"", "", commandName}))
}