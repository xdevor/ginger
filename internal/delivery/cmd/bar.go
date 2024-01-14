package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(barCmd)
}

var barCmd = &cobra.Command{
	Use:   "bar",
	Short: "Bar command is a demo command",
	Long:  `Command bar is a demo command. ref: https://cobra.dev/`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bar command power by cobra")
	},
}
