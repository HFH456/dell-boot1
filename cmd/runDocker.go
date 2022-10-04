package cmd

import (
	"dell-boot1/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var runDockerCmd = &cobra.Command{
	Use:   "run",
	Short: "mytool run docker operations",
	Long:  `mytool run: run docker by docker-compose/yml or docker.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("You don't need to add any args!")
			return
		}
		composeFilePaths := "./docker-compose.yml"

		compose := utils.DockerCompose{}
		compose.GetConf(composeFilePaths)

		err := utils.StartDockerByYml(compose)
		if err != nil {
			fmt.Printf("Could not run compose file: %v - %v", composeFilePaths, err)
			return
		}
	},
}

func initRun() {
	rootCmd.AddCommand(runDockerCmd)
}
