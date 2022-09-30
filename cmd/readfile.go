package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var readYmlCmd = &cobra.Command{
	Use:   "readyml",
	Short: "mytool name get operations",
	Long:  `mytool get: get name info in names dir`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Only one name can be get in one time! Or use get all to get all names!")
			return
		}
		fileDir := args[0]
		ymlInfo, err := ioutil.ReadFile(fileDir)
		if err != nil {
			fmt.Printf("%s not found\n", fileDir)
			return
		}
		fmt.Println(fileDir, ":", string(ymlInfo))
	},
}

func initRead() {
	rootCmd.AddCommand(readYmlCmd)
}
