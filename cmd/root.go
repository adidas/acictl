package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute command tool
func Execute() {
	rootCmd.AddCommand(
		initCmd,
		configCmd,
		deployCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func root(cmd *cobra.Command, args []string) {
	if version {
		fmt.Println(cmd.Long)
	} else {
		cmd.Help()
	}
}
