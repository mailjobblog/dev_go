// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wire-example/internal/config"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	string2 := _wireStringValue
	configString2 := _wireString2Value
	configConfig, err := config.New(string2, configString2)
	if err != nil {
		return nil, err
	}
	app := NewApp(configConfig)
	return app, nil
}

var (
	_wireStringValue  = "demo string1"
	_wireString2Value = config.String2("demo string 2")
)
