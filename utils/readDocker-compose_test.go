package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadYml(t *testing.T) {
	services := DockerCompose{
		Version: "3",
		Services: map[string]ContainerItem{
			"init-nginx": ContainerItem{
				"nginx",
				"init-nginx",
				[]string{
					"80",
				},
				true,
			},
		},
	}

	dir := "../yml-test.yml"
	compose := DockerCompose{}
	compose.GetConf(dir)

	a := assert.New(t)
	a.Equal(services, compose)
}
