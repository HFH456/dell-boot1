package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func (c *DockerCompose) GetConf(dir string) (*DockerCompose, error) {
	yamlFile, err := ioutil.ReadFile(dir)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c, nil
}

type DockerCompose struct {
	Version  string                   `yaml:"version"`
	Services map[string]ContainerItem `yaml:"services"`
}
type ContainerItem struct {
	Image         string   `yaml:"image"`
	ContainerName string   `yaml:"container_name"`
	Ports         []string `yaml:"ports"`
	Init          bool     `yaml:"init_container"`
}
