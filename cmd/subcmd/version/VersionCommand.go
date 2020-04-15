package version

import (
	"fmt"
	"fops/cmd/subcmd"
	"github.com/spf13/cobra"
)
var version string = "v0.0.1"
var usage string = "version"
var shortName string = "Show the version info"
type versionCmd struct {
	subcmd.SubCommand
}
func GetSubCommand() (*cobra.Command, error){
	return subcmd.GetCommand(&versionCmd{subcmd.SubCommand{
		Usage:     usage,
		ShortName: shortName,
	}})
}
func (c *versionCmd) SetRequired(cmd *cobra.Command) error {
	return nil
}
func (c *versionCmd) SetFlags(cmd *cobra.Command) error {
	return nil
}
func (c *versionCmd) RUN (cmd *cobra.Command, args []string) error {
	fmt.Println(fmt.Sprintf("fops %s", version))
	return nil
}