package sha

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetAlgorithm(t *testing.T) {
	assert.Equal(t, "*sha.algorithmSHA", reflect.TypeOf(GetAlgorithm()).String())
}

func TestGetHash(t *testing.T) {
	alg := GetAlgorithm()
	assert.Equal(t, "string", reflect.TypeOf(alg.GetHash([]byte{})).String())
}