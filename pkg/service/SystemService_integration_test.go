// +build integration

package service_test

import (
	"fops/pkg/service"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInitialSystemService(t *testing.T) {
	assert.NoError(t, service.InitialSystemService())
}

func TestGetSystemService(t *testing.T) {
	assert.Equal(t, "*service.SystemService", reflect.TypeOf(service.GetSystemService()).String())
}

func TestAddPreload(t *testing.T) {
	assert.NoError(t, (&service.SystemService{}).AddPreload(func(){}))
}

func TestShutdown(t *testing.T) {
	system := &service.SystemService{}
	assert.NoError(t, system.AddPreload(func() {}))
	assert.NotPanics(t, func() {
		system.Shutdown()
	})
}
