// +build integration

package md5_test

import (
	"fops/cmd/subcmd/checksum/algorithm/md5"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetAlgorithm(t *testing.T) {
	assert.Equal(t, "*md5.algorithmMD5", reflect.TypeOf(md5.GetAlgorithm()).String())
}

func TestGetHash(t *testing.T) {
	alg := md5.GetAlgorithm()
	assert.Equal(t, "string", reflect.TypeOf(alg.GetHash([]byte{})).String())
}