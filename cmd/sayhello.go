package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sayhelloCmd represents the sayhello command
var sayhelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}
		fmt.Println("Hello " + name)
	},
}

func init() {
	sayCmd.AddCommand(sayhelloCmd)
	sayhelloCmd.Flags().StringP("name", "n", "", "Set your name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sayhelloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sayhelloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
