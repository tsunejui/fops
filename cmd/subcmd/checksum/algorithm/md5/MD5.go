package md5

import (
	"crypto/sha1"
	"fmt"
	"fops/cmd/subcmd/checksum/algorithm"
)

type algorithmMD5 struct {
	algorithm.Algorithm
}

var algorithmName string = "md5"

func GetAlgorithm() algorithm.InterfaceAlgorithm {
	return &algorithmMD5{
		algorithm.Algorithm{
			Name: algorithmName,
		},
	}
}

func (a *algorithmMD5) GetHash(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}
