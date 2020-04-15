package root

import (
	"github.com/spf13/cobra"
)
//type RootCommand struct {
//	Flags *getopt.FlagSet
//}
//var rootCommand *RootCommand
//var help bool

var rootCmd *cobra.Command

func SetRootCommand(cmd *cobra.Command) error {
	rootCmd = cmd
	return nil
}

func GetRootCommand() (*cobra.Command, error) {
	return rootCmd, nil
}

//func InitialRootCommand() error {
//	flags := getopt.NewFlagSet("main", flag.ExitOnError)
//	flags.BoolVar(&help, "h", false, "Help about commands")
//	flags.Alias("h", "help")
//	rootCommand = &RootCommand {
//		Flags: flags,
//	}
//	return nil
//}
//func GetRootCommand () *RootCommand{
//	return rootCommand
//}
//func (r RootCommand) Do(args []string) error {
//	if parseErr := r.Flags.Parse(args[1:]); parseErr != nil {
//		return errors.New("Failed to parse root command")
//	}
//	if help {
//		return errors.New("Show command help")
//	}
//	return errors.New("nothing to do")
//	//return nil
//}
