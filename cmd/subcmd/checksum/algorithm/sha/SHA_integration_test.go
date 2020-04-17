// +build integration

package sha_test

import (
	"fops/cmd/subcmd/checksum/algorithm/sha"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetAlgorithm(t *testing.T) {
	assert.Equal(t, "*sha.algorithmSHA", reflect.TypeOf(sha.GetAlgorithm()).String())
}

func TestGetHash(t *testing.T) {
	alg := sha.GetAlgorithm()
	assert.Equal(t, "string", reflect.TypeOf(alg.GetHash([]byte{})).String())
}
