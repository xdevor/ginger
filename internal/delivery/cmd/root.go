package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "foo command is a demo command",
	Long:  `Command foo is a demo command. ref: https://cobra.dev/`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Foo command power by cobra")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
