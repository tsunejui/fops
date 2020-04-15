package subcmd

import (
	"github.com/prashantv/gostub"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCommand struct {
	InterfaceCommand
	mock.Mock
}

func (m *mockCommand) SetRequired(cmd *cobra.Command) error {
	mockArgs := m.Called(cmd)
	return mockArgs.Error(0)
}
func (m *mockCommand) SetFlags(cmd *cobra.Command) error {
	mockArgs := m.Called(cmd)
	return mockArgs.Error(0)
}
func (m *mockCommand) GetShortName() string {
	mockArgs := m.Called()
	return mockArgs.Get(0).(string)
}
func (m *mockCommand) GetUsage() string {
	mockArgs := m.Called()
	return mockArgs.Get(0).(string)
}
func (m *mockCommand) RUN (cmd *cobra.Command, args []string) error {
	mockArgs := m.Called(cmd, args)
	return mockArgs.Error(0)
}
func TestSubCommandGetShortName(t *testing.T) {
	var shortName string = "test"
	assert.Equal(t, (&SubCommand{
		ShortName: shortName,
		Usage:     "",
	}).GetShortName(), shortName)
}
func TestSubCommandGetUsage(t *testing.T) {
	var usage string = "test"
	assert.Equal(t, (&SubCommand{
		ShortName: "",
		Usage:     usage,
	}).GetUsage(), usage)
}

func TestGetCommand(t *testing.T) {
	mockCmd := &mockCommand{}
	mockUsage := "test"
	mockShort := "test"
	mockSubCmd := &cobra.Command{
		Use:   mockUsage,
		Short: mockShort,
	}
	defer gostub.StubFunc(&getCobraCommand, mockSubCmd).Reset()
	mockCmd.On("SetRequired", mockSubCmd).Return(nil)
	mockCmd.On("SetFlags", mockSubCmd).Return(nil)
	cmd, cmdErr := GetCommand(mockCmd)
	assert.NoError(t, cmdErr)
	assert.Equal(t, cmd, mockSubCmd)
}