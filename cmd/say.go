package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Provide item to the say command")
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
