package main

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	fileName := fmt.Sprintf("./%s_config.json", env)

	gonfig.GetConf(fileName, &configuration)

	return configuration
}
