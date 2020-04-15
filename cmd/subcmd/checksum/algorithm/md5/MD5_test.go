package md5

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetAlgorithm(t *testing.T) {
	assert.Equal(t, "*md5.algorithmMD5", reflect.TypeOf(GetAlgorithm()).String())
}

func TestGetHash(t *testing.T) {
	alg := GetAlgorithm()
	assert.Equal(t, "string", reflect.TypeOf(alg.GetHash([]byte{})).String())
}