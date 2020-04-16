package service

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInitialSystemService(t *testing.T) {
	assert.NoError(t, InitialSystemService())
}

func TestGetSystemService(t *testing.T) {
	assert.Equal(t, "*service.SystemService", reflect.TypeOf(GetSystemService()).String())
}

func TestAddPreload(t *testing.T) {
	assert.NoError(t, (&SystemService{}).AddPreload(func(){}))
}

func TestShutdown(t *testing.T) {
	system := &SystemService{}
	assert.NoError(t, system.AddPreload(func() {}))
	assert.NotPanics(t, func() {
		system.Shutdown()
	})
}