package checksum

import (
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"fops/commands"
	"fops/service"
	"io/ioutil"
	"rsc.io/getopt"
	"crypto/md5"
)

type command struct {
	commands.SubCommand
	path string
	md5Algorithm bool
	sha1Algorithm bool
	sha256Algorithm bool
}
func (c *command) Initial () error {
	c.GetField().StringVar(&c.path, "file", "", " the input file")
	c.GetField().Alias("f", "file")
	c.GetField().BoolVar(&c.md5Algorithm,"md5",false, "generate md5 checksum")
	c.GetField().BoolVar(&c.sha1Algorithm,"sha1",false, "generate sha1 checksum")
	c.GetField().BoolVar(&c.sha256Algorithm,"sha256",false, "generate sha256 checksum")
	return nil
}
func (c *command) Handle (args []string) error {
	if len(c.path) > 0 {
		data, getDataErr := getFileData(c.path)
		if getDataErr != nil {
			fmt.Println(getDataErr)
			return nil
		}
		switch true {
		case c.md5Algorithm:
			fmt.Println(fmt.Sprintf("%x", md5.Sum(data)))
			break
		case c.sha1Algorithm:
			fmt.Println(fmt.Sprintf("%x", sha1.Sum(data)))
			break
		case c.sha256Algorithm:
			fmt.Println(fmt.Sprintf("%x", sha256.Sum256(data)))
			break
		default:
			return errors.New("Invalid Argument")
		}
	}else{
		return errors.New("Missing Argument")
	}
	return nil
}
func GetCommand() commands.CommandInterface {
	checksum := getopt.NewFlagSet("checksum", flag.ExitOnError)
	return &command{
		SubCommand: commands.SubCommand{
			Flag: checksum,
			CommandName: "checksum",
			Description: "Print checksum of file",
		},
	}
}

var getFileData = func (path string) ([]byte, error){
	_, _, fileErr := service.GetFile(path)
	if fileErr != nil {
		return []byte{}, fileErr
	}
	data, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return []byte{}, errors.New("Failed to read file: " + readErr.Error())
	}
	return data, nil
}