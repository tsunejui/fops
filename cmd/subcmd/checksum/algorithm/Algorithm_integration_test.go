// +build integration

package algorithm_test

import (
	"fops/cmd/subcmd/checksum/algorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestGetName(t *testing.T) {
	var mockName string = "test"
	assert.Equal(t, (&algorithm.Algorithm{Name:mockName}).GetName(), mockName)
}