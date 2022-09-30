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
		if len(args) != 0 {
			fmt.Println("You don't need to add any args!")
			return
		}
		composeFilePaths := []string{"./docker-compose.yml"}

		err := utils.RunDockerCompose()
		if err != nil {
			fmt.Printf("Could not run compose file: %v - %v", composeFilePaths, err)
			return
		}
	},
}

func initDown() {
	rootCmd.AddCommand(runDockerCmd)
}