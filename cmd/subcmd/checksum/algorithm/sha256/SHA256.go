package sha256

import (
	"crypto/sha256"
	"fmt"
	"fops/cmd/subcmd/checksum/algorithm"
)

type algorithmSHA256 struct {
	algorithm.Algorithm
}
var algorithmName string = "sha256"
func GetAlgorithm() algorithm.InterfaceAlgorithm {
	return &algorithmSHA256{
		algorithm.Algorithm{
			Name: algorithmName,
		},
	}
}

func (a *algorithmSHA256) GetHash(data []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(data))
}
