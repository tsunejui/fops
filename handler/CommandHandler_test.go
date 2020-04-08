package handler

import (
	"errors"
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

func (m *mockSubCommand) Parse(parseArgs []string) error {
	args := m.Called(parseArgs)
	return args.Error(0)
}

func TestInitialCommandHandler(t *testing.T) {
	assert.NoError(t, InitialCommandHandler())
}

func TestGetCommandHandler(t *testing.T) {
	assert.Equal(t, "*handler.CommandHandler", reflect.TypeOf(GetCommandHandler()).String())
}

func TestAddCommand(t *testing.T) {
	assert.NotPanics(t, func() {
		(&CommandHandler{}).AddCommand(&mockSubCommand{
			SubCommand: commands.SubCommand{
				Flag: getopt.NewFlagSet("test", flag.ExitOnError),
				Description: "",
				CommandName: "",
			},
		})
	})
}

func TestGetCommands(t *testing.T) {
	mockCommands := []commands.CommandInterface{
		&mockSubCommand{
			SubCommand: commands.SubCommand{
				Flag: getopt.NewFlagSet("test", flag.ExitOnError),
				Description: "",
				CommandName: "",
			},
		},
	}
	assert.Equal(t, mockCommands, (&CommandHandler{
		Commands:    mockCommands,
		RootCommand: nil,
	}).GetCommands())
}

func TestDoWithInitialError(t *testing.T) {
	defer gostub.StubFunc(&getRootCommand, nil).Reset()
	command := &mockSubCommand{
		SubCommand: commands.SubCommand{
			Flag: getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "test",
		},
	}
	command.On("Initial").Return(errors.New("test"))
	mockCommands := []commands.CommandInterface{command}
	assert.NoError(t, (&CommandHandler{
		Commands:    mockCommands,
		RootCommand: nil,
	}).Do([]string{"", "test"}))
}

func TestDoWithParseError(t *testing.T) {
	args := []string{"", "test", ""}
	defer gostub.StubFunc(&getRootCommand, nil).Reset()
	command := &mockSubCommand{
		SubCommand: commands.SubCommand{
			Flag: getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "test",
		},
	}
	command.On("Initial").Return(nil)
	command.On("Parse", args[2:]).Return(errors.New("test"))
	mockCommands := []commands.CommandInterface{command}
	assert.NoError(t, (&CommandHandler{
		Commands:    mockCommands,
		RootCommand: nil,
	}).Do(args))
}

func TestDoWithHandleError(t *testing.T) {
	args := []string{"", "test", ""}
	defer gostub.StubFunc(&getRootCommand, nil).Reset()
	command := &mockSubCommand{
		SubCommand: commands.SubCommand{
			Flag: getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "test",
		},
	}
	command.On("Initial").Return(nil)
	command.On("Parse", args[2:]).Return(nil)
	command.On("Handle", args).Return(errors.New("test"))
	mockCommands := []commands.CommandInterface{command}
	assert.NoError(t, (&CommandHandler{
		Commands:    mockCommands,
		RootCommand: nil,
	}).Do(args))
}

func TestDoWithRootCommandError(t *testing.T) {
	args := []string{"", "test", ""}
	defer gostub.StubFunc(&getRootCommand, errors.New("test")).Reset()
	command := &mockSubCommand{
		SubCommand: commands.SubCommand{
			Flag: getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "test",
		},
	}
	command.On("Initial").Return(nil)
	command.On("Parse", args[2:]).Return(nil)
	command.On("Handle", args).Return(nil)
	mockCommands := []commands.CommandInterface{command}
	assert.NoError(t, (&CommandHandler{
		Commands:    mockCommands,
		RootCommand: nil,
	}).Do(args))
}

func TestDoWithNoError(t *testing.T) {
	args := []string{"", "test", ""}
	defer gostub.StubFunc(&getRootCommand, nil).Reset()
	command := &mockSubCommand{
		SubCommand: commands.SubCommand{
			Flag: getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "test",
		},
	}
	command.On("Initial").Return(nil)
	command.On("Parse", args[2:]).Return(nil)
	command.On("Handle", args).Return(nil)
	mockCommands := []commands.CommandInterface{command}
	assert.NoError(t, (&CommandHandler{
		Commands:    mockCommands,
		RootCommand: nil,
	}).Do(args))
}