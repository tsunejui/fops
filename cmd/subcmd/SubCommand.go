package subcmd

import (
	"github.com/spf13/cobra"
)

type InterfaceCommand interface {
	SetRequired(cmd *cobra.Command) error
	SetFlags(cmd *cobra.Command) error
	GetShortName() string
	GetUsage() string
	RUN (cmd *cobra.Command, args []string) error
}

type SubCommand struct {
	ShortName string
	Usage string
}

var getCobraCommand = func(cmd InterfaceCommand) *cobra.Command{
	return &cobra.Command{
		Use:   cmd.GetUsage(),
		Short: cmd.GetShortName(),
	}
}

func (c *SubCommand) GetShortName() string {
	return c.ShortName
}
func (c *SubCommand) GetUsage() string {
	return c.Usage
}

func GetCommand(subCmd InterfaceCommand) (*cobra.Command, error){
	cmd := getCobraCommand(subCmd)
	if flgErr := subCmd.SetFlags(cmd); flgErr != nil {
		return nil, flgErr
	}
	if reqErr := subCmd.SetRequired(cmd); reqErr != nil {
		return nil, reqErr
	}
	cmd.RunE = subCmd.RUN
	return cmd, nil
}

