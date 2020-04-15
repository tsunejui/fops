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

func (c *SubCommand) GetShortName() string {
	return c.ShortName
}
func (c *SubCommand) GetUsage() string {
	return c.Usage
}

func GetCommand(subCmd InterfaceCommand) (*cobra.Command, error){
	cmd := &cobra.Command{
		Use:   subCmd.GetUsage(),
		Short: subCmd.GetShortName(),
	}
	if flgErr := subCmd.SetFlags(cmd); flgErr != nil {
		return nil, flgErr
	}
	if reqErr := subCmd.SetRequired(cmd); reqErr != nil {
		return nil, reqErr
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if err := subCmd.RUN(cmd, args); err != nil {
			//fmt.Println(cmd.Usage())
		}
	}
	return cmd, nil
}
