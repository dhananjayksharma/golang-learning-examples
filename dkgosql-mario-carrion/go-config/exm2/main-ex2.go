package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"path"

	yaml "gopkg.in/yaml.v3"
)

//go:embed config/*.yaml
var content embed.FS

func main() {
	var env string
	flag.StringVar(&env, "env", "dev", "environment to use")
	flag.Parse()

	content, err := content.Open(path.Join("config", fmt.Sprintf("%s.yaml", env)))
	if err != nil {
		log.Fatalln("content.Open", err)
	}
	config := struct {
		Credentials struct {
			Password string `yaml:"password"`
		}
	}{}

	decoder := yaml.NewDecoder(content)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("decoder.Decode:", err)
	}

	fmt.Printf("password: %+v\n", config.Credentials.Password)
}
