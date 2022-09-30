package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mytool",
	Short: "mytool is a tool to record names", // 子命令的简单说明，务必简短，因为会出现在其上级命令的help结果中，过长会导致换行、不美观，此外还会出现在自动补全时，过长容易影响视线
	Long: `The best tool to record names in the world! // 子命令的功用详细说明，可以写的比较详细些
Just try it!!!
Complete documentation is available at http://mytool.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use mytool -h or --help for help.")
	},
}

func initAll() {
	initRun()
	initDown()
}

func Execute() {
	initAll()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
