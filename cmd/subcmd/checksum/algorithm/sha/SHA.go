package sha

import (
	"crypto/sha1"
	"fmt"
	"fops/cmd/subcmd/checksum/algorithm"
)

type algorithmSHA struct {
	algorithm.Algorithm
}
var algorithmName string = "sha1"
func GetAlgorithm() algorithm.InterfaceAlgorithm {
	return &algorithmSHA{
		algorithm.Algorithm{
			Name: algorithmName,
		},
	}
}

func (a *algorithmSHA) GetHash(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}
