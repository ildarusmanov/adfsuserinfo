package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var configYamlPath = "./test/fixtures/config.yml"

func TestLoadConfigYAML(t *testing.T) {
    config, err := LoadConfigYAML(configYamlPath)

    assert := assert.New(t)
    assert.Nil(err)
    assert.NotNil(config)
    assert.NotNil(config.ServerAddr)
    assert.NotNil(config.ServerHost)
}
