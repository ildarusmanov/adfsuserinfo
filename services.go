package main

import (
	"github.com/ildarusmanov/adfsuserinfo/services"
)

func BuildServiceLocator(config *Config) (*services.ServiceLocator, error) {
	userinfo := services.CreateUserinfoService()

	locator := services.CreateServiceLocator()
	locator.Set("Userinfo", userinfo)

	return locator, nil
}
