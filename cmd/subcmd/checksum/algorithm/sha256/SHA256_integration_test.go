// +build integration

package sha256_test

import (
	"fops/cmd/subcmd/checksum/algorithm/sha256"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetAlgorithm(t *testing.T) {

	assert.Equal(t, "*sha256.algorithmSHA256", reflect.TypeOf(sha256.GetAlgorithm()).String())
}

func TestGetHash(t *testing.T) {
	alg := sha256.GetAlgorithm()
	assert.Equal(t, "string", reflect.TypeOf(alg.GetHash([]byte{})).String())
}
