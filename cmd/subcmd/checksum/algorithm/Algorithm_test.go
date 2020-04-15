package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetName(t *testing.T) {
	var mockName string = "test"
	assert.Equal(t, (&Algorithm{Name:mockName}).GetName(), mockName)
}
