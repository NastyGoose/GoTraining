package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Executes passed command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(strings.Join(args, " "))

		out, err := exec.Command("sh", "-c", strings.Join(args, " ")).Output()

		if err != nil {
			fmt.Printf("%s", err)

		}
		fmt.Printf("%s", out)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}
