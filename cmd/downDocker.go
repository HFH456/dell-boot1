package cmd

import (
	"dell-boot1/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var downDockerCmd = &cobra.Command{
	Use:   "down",
	Short: "mytool down docker operations",
	Long:  `mytool run: run docker by docker-compose/yml or docker.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You have to input the identifier of existing running compose!")
			return
		}
		composeFilePaths := []string{"./docker-compose.yml"}
		identifier := args[0]

		err := utils.DownDockerCompose(identifier)
		if err != nil {
			fmt.Printf("Could not run compose file: %v - %v", composeFilePaths, err)
			return
		}
	},
}

func initDown() {
	rootCmd.AddCommand(downDockerCmd)
}
