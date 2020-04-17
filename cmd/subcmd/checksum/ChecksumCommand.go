package checksum

import (
	"errors"
	"fmt"
	"fops/cmd/subcmd"
	"fops/cmd/subcmd/checksum/algorithm"
	md5 "fops/cmd/subcmd/checksum/algorithm/md5"
	sha "fops/cmd/subcmd/checksum/algorithm/sha"
	sha256 "fops/cmd/subcmd/checksum/algorithm/sha256"
	"fops/pkg/service"
	"github.com/spf13/cobra"
	"io/ioutil"
)
var usage string = "checksum"
var shortName string = "Print checksum of file"
var algorithmList []algorithm.InterfaceAlgorithm
var getFileArgs = func(cmd *cobra.Command) (string, error){
	return cmd.Flags().GetString("file")
}
var getFlagName = func(cmd *cobra.Command, name string) (bool, error){
	return cmd.Flags().GetBool(name)
}
var getFileData = func (path string) ([]byte, error){
	_, _, fileErr := service.GetFile(path)
	if fileErr != nil {
		return []byte{}, fileErr
	}
	data, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return []byte{}, fmt.Errorf("Failed to read file:  %s", readErr.Error())
	}
	return data, nil
}
var markFlag = func(cmd *cobra.Command, name string) error {
	return cmd.MarkFlagRequired(name)
}
type checksum struct {
	subcmd.SubCommand
}
func GetSubCommand() (*cobra.Command, error){
	if initErr := initialAlgorithm(); initErr != nil {
		return nil, initErr
	}
	return subcmd.GetCommand(&checksum{subcmd.SubCommand{
		Usage:     usage,
		ShortName: shortName,
	}})
}

func (c *checksum) SetRequired(cmd *cobra.Command) error {
	if markErr := markFlag(cmd, "file"); markErr != nil {
		return markErr
	}
	return nil
}
func (c *checksum) SetFlags(cmd *cobra.Command) error {
	cmd.Flags().StringP("file", "f", "", "the input file")
	for _, alg := range algorithmList {
		cmd.Flags().Bool(alg.GetName(), false, fmt.Sprintf("generate %s checksum", alg.GetName()))
	}
	return nil
}
func (c *checksum) RUN (cmd *cobra.Command, args []string) error {
	filepath, pathErr := getFileArgs(cmd)
	if pathErr != nil {
		return pathErr
	}
	if len(filepath) > 0 {
		data, getDataErr := getFileData(filepath)
		if getDataErr != nil {
			return getDataErr
		}
		var parse bool
		loop:
			for _, alg := range algorithmList {
				if ok, _ := getFlagName(cmd, alg.GetName()); ok {
					fmt.Println(alg.GetHash(data))
					parse = true
					break loop
				}
			}
		if !parse {
			return errors.New("Invalid Argument")
		}
	}else{
		return errors.New("Missing Argument")
	}
	return nil
}
func initialAlgorithm() error {
	algorithmList = []algorithm.InterfaceAlgorithm{
		md5.GetAlgorithm(),
		sha.GetAlgorithm(),
		sha256.GetAlgorithm(),
	}
	return nil
}