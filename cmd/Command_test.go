package cmd

import (
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecute(t *testing.T) {
	defer gostub.StubFunc(&executeCommand, nil).Reset()
	assert.NoError(t, Execute())
}