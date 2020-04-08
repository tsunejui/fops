package templates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelpCommandUsage(t *testing.T) {
	assert.NotPanics(t, func() {
		HelpCommandUsage("test")
	})
}

func TestSubCommandUsage(t *testing.T) {
	assert.NotPanics(t, func() {
		SubCommandUsage("test", "test")
	})
}

func TestRootCommandUsage(t *testing.T) {
	assert.NotPanics(t, func() {
		RootCommandUsage([]CommandTemplate{
			{
				Name:        "test",
				Description: "test",
			},
		})
	})
}