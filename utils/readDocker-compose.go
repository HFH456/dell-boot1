package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type conf struct {
	Version  string `yaml:"version"`
	Services struct {
	} `yaml:"services"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("docker-compose.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
