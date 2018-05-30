package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildServiceLocator(t *testing.T) {
	config := CreateNewConfig()

	locator, err := BuildServiceLocator(config)

	assert.Nil(t, err)
	assert.NotNil(t, locator)
}
