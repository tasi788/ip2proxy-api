package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func ConfigParser() Config {
	file, err := os.Open("./config.yml")
	if err != nil {
		log.Panic("Can Not Load config.yml")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	ConfigByte, _ := ioutil.ReadAll(file)

	t := Config{}

	loadErr := yaml.Unmarshal(ConfigByte, &t)
	if loadErr != nil {
		log.Panic("Parse config.yml Error")
	}
	return t
}
