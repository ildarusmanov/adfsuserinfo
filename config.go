package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
    ServerHost      string `yaml:"server_host"`
    ServerAddr      string `yaml:"server_addr"`
}

// constructor
func CreateNewConfig() *Config {
    return &Config{}
}

// load config from YAML file
func LoadConfigYAML(configFilePath string) (*Config, error) {
    config := CreateNewConfig()

    configFileData, err := ioutil.ReadFile(configFilePath)

    if err != nil {
        return nil, err
    }

    if err = yaml.Unmarshal([]byte(configFileData), config); err != nil {
        return nil, err
    }

    return config, nil
}
