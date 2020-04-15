package version

import (
	"fmt"
	"fops/cmd/subcmd"
	"github.com/spf13/cobra"
)
var version string = "v0.0.1"
var usage string = "version"
var shortName string = "Show the version info"
type checksum struct {
	subcmd.SubCommand
}
func GetSubCommand() (*cobra.Command, error){
	return subcmd.GetCommand(&checksum{subcmd.SubCommand{
		Usage:     usage,
		ShortName: shortName,
	}})
}
func (c *checksum) SetRequired(cmd *cobra.Command) error {
	return nil
}
func (c *checksum) SetFlags(cmd *cobra.Command) error {
	return nil
}
func (c *checksum) RUN (cmd *cobra.Command, args []string) error {
	fmt.Println(fmt.Sprintf("fops %s", version))
	return nil
}